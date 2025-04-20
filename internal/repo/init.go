package repo

import (
	"context"
	"database/sql"
	"fmt"
	"master-finanacial-planner/internal/entity"
	"master-finanacial-planner/internal/logger"
)

type ResourceRepo interface {
	GetAssetClass(ctx context.Context) ([]entity.AssetClass, error)
	GetAllAllocationTypeConfig(ctx context.Context) ([]AllocationTypeConfig, error)
	GetInvestingSurplus(ctx context.Context) (float64, error)
	GetLiquidAndIlliquidAssets(ctx context.Context) (map[string]float64, error)
	GetAllLiability(ctx context.Context) (float64, error)
}

type ResourceRepository struct {
	db *sql.DB
}

func NewResource(db *sql.DB) *ResourceRepository {
	return &ResourceRepository{
		db,
	}
}

func (r *ResourceRepository) GetAssetClass(ctx context.Context) ([]entity.AssetClass, error) {
	var assetClasses []entity.AssetClass

	query := "SELECT id, name, expected_return_in_percentage FROM asset_class"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error querying asset class data: %v", err)
	}
	defer rows.Close() // Ensure the rows iterator is closed properly

	for rows.Next() {
		var assetClass entity.AssetClass
		if err := rows.Scan(&assetClass.ID, &assetClass.Name, &assetClass.ExpectedReturnInPercentage); //&assetClass.ExpectedReturnInPercentage
		err != nil {
			logger.LogError(ctx, err.Error())
			return nil, fmt.Errorf("error scanning asset class row: %v", err)
		}
		assetClasses = append(assetClasses, assetClass)
	}

	// Check for any errors that might have occurred during iteration
	if err := rows.Err(); err != nil {

		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	// Return the slice of AssetClass
	return assetClasses, nil
}

type AllocationTypeConfig struct {
	ID                     int64   `json:"id"`                            // bigint corresponds to int64 in Go
	AllocationTypeName     string  `json:"allocation_type_name"`          // varchar corresponds to string
	AssetReturns           float64 `json:"expected_return_in_percentage"` // varchar corresponds to string
	AllocationTypeId       int64   `json:"allocation_type_id"`            // varchar corresponds to string
	AssetClassID           float64 `json:"asset_class_id"`                // double precision corresponds to float64
	AllocationInPercentage float64 `json:"allocation_in_percentage"`      // double precision corresponds to float64
}

func (r *ResourceRepository) GetAllAllocationTypeConfig(ctx context.Context) ([]AllocationTypeConfig, error) {
	var assetClasses []AllocationTypeConfig

	query := `SELECT 
				atc.id, 
				at.name as allocation_type_name, 
				ac.expected_return_in_percentage , 
				atc.allocation_type_id, 
				atc.asset_class_id, 
				atc.allocation_in_percentage 
			FROM allocation_type_config atc
			JOIN allocation_type at
				ON atc.allocation_type_id = at.id
			JOIN asset_class ac
				ON atc.asset_class_id = ac.id`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error querying asset class data: %v", err)
	}
	defer rows.Close() // Ensure the rows iterator is closed properly

	for rows.Next() {
		var assetClass AllocationTypeConfig
		if err := rows.Scan(
			&assetClass.ID,
			&assetClass.AllocationTypeName,
			&assetClass.AssetReturns,
			&assetClass.AllocationTypeId,
			&assetClass.AssetClassID,
			&assetClass.AllocationInPercentage,
		); err != nil {
			logger.LogError(ctx, err.Error())
			return nil, fmt.Errorf("error scanning class row: %v", err)
		}
		assetClasses = append(assetClasses, assetClass)
	}

	// Check for any errors that might have occurred during iteration
	if err := rows.Err(); err != nil {

		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	// Return the slice of AssetClass
	return assetClasses, nil
}

func (r *ResourceRepository) GetInvestingSurplus(ctx context.Context) (float64, error) {
	query := `SELECT 
					SUM(CASE
							WHEN is_inflow = TRUE THEN amount
							WHEN is_inflow = FALSE THEN -amount
						END) AS total_surplus
				FROM cashflow;`

	var totalSurplus float64

	// Use QueryRowContext for a query expecting a single row of data
	err := r.db.QueryRowContext(ctx, query).Scan(&totalSurplus)
	if err != nil {
		// Log the error and return a detailed error message
		logger.LogError(ctx, fmt.Sprintf("error querying investing surplus: %v", err))
		return 0, fmt.Errorf("error querying investing surplus: %w", err)
	}

	// Return the total surplus
	return totalSurplus, nil
}

func (r *ResourceRepository) GetLiquidAndIlliquidAssets(ctx context.Context) (map[string]float64, error) {
	// Define the query to get liquid and illiquid asset sums
	query := `
		SELECT 
			type,
			SUM(amount) 
		FROM investments
		GROUP BY type
	`

	// Prepare a map to store the results
	assets := make(map[string]float64)

	// Execute the query
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error querying liquid and illiquid assets data: %v", err)
	}
	defer rows.Close()

	// Iterate through the results
	for rows.Next() {
		var assetType string
		var totalAmount float64

		// Scan the row into variables
		if err := rows.Scan(&assetType, &totalAmount); err != nil {
			logger.LogError(ctx, err.Error())
			return nil, fmt.Errorf("error scanning liquid and illiquid asset row: %v", err)
		}

		// Store the result in the map
		assets[assetType] = totalAmount
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	return assets, nil
}

func (r *ResourceRepository) GetAllLiability(ctx context.Context) (float64, error) {
	// Define the query to get the sum of liabilities
	query := `SELECT SUM(amount) FROM liabilities`

	var totalAmount float64

	// Use QueryRowContext since we expect a single value
	err := r.db.QueryRowContext(ctx, query).Scan(&totalAmount)
	if err != nil {
		logger.LogError(ctx, fmt.Sprintf("Error querying total liabilities: %v. Query: %s", err, query))
		return 0, fmt.Errorf("error querying total liabilities: %w", err)
	}

	return totalAmount, nil
}
