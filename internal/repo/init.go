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
	GetGoals(ctx context.Context) ([]entity.Goals, error)
	GetAllocationByYearLeft(ctx context.Context, yearsLeft int64) ([]entity.AllocationType, error)
	GetAllocationConfigByAllocationTypeId(ctx context.Context, allocationTypeId int64) ([]entity.AllocationTypeConfig, error)
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

func (r *ResourceRepository) GetGoals(ctx context.Context) ([]entity.Goals, error) {
	var goals []entity.Goals

	query := `SELECT 
				id,
				name,
				description,
				years_left,
				inflation_percentage,
				today_amount,
				allocated_amount,
				sip_step_up_percentage
			  FROM goals`

	// Execute the query
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error querying goals data: %v", err)
	}
	defer rows.Close() // Ensure the rows iterator is closed properly

	// Iterate through the rows
	for rows.Next() {
		var goal entity.Goals
		if err := rows.Scan(
			&goal.ID,
			&goal.Name,
			&goal.Description,
			&goal.YearsLeft,
			&goal.InflationPercentage,
			&goal.TodayAmount,
			&goal.AllocatedAmount,
			&goal.SIPStepUpPercentage,
		); err != nil {
			logger.LogError(ctx, err.Error())
			return nil, fmt.Errorf("error scanning goal row: %v", err)
		}
		goals = append(goals, goal)
	}

	// Check for any errors that occurred during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	// Return the list of goals
	return goals, nil
}

func (r *ResourceRepository) GetAllocationByYearLeft(ctx context.Context, yearsLeft int64) ([]entity.AllocationType, error) {
	var allocationTypes []entity.AllocationType

	query := `SELECT 
					id,
					name
			  FROM 
					allocation_type
			  WHERE 
					$1 >= min_age 
					AND $1 <= COALESCE(max_age, $1);`

	// Execute the query with the yearsLeft parameter
	rows, err := r.db.QueryContext(ctx, query, yearsLeft)
	if err != nil {
		return nil, fmt.Errorf("error querying allocation types for years left: %v", err)
	}
	defer rows.Close() // Ensure the rows iterator is closed properly

	// Iterate through the rows and scan results into the struct
	for rows.Next() {
		var allocationType entity.AllocationType
		if err := rows.Scan(
			&allocationType.ID,
			&allocationType.Name,
		); err != nil {
			logger.LogError(ctx, err.Error())
			return nil, fmt.Errorf("error scanning allocation type row: %v", err)
		}
		allocationTypes = append(allocationTypes, allocationType)
	}

	// Check for any errors that occurred during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	// Return the list of allocation types
	return allocationTypes, nil
}

func (r *ResourceRepository) GetAllocationConfigByAllocationTypeId(ctx context.Context, allocationTypeId int64) ([]entity.AllocationTypeConfig, error) {
	var allocationTypeConfigs []entity.AllocationTypeConfig

	query := `SELECT  ac.name,
				COALESCE(atc.allocation_in_percentage, 0) as allocation_in_percentage
				FROM allocation_type_config AS atc
				RIGHT OUTER JOIN asset_class AS ac
					ON atc.asset_class_id = ac.id AND atc.allocation_type_id = $1;`

	// Execute the query with the yearsLeft parameter
	rows, err := r.db.QueryContext(ctx, query, allocationTypeId)
	if err != nil {
		return nil, fmt.Errorf("error querying allocation types config for allocation type: %v", err)
	}
	defer rows.Close() // Ensure the rows iterator is closed properly

	// Iterate through the rows and scan results into the struct
	for rows.Next() {
		var allocationTypeConfig entity.AllocationTypeConfig
		if err := rows.Scan(
			&allocationTypeConfig.AssetName,
			&allocationTypeConfig.AllocationInPercentage,
			//&allocationTypeConfig.AssetReturnInPercentage,
		); err != nil {
			logger.LogError(ctx, err.Error())
			return nil, fmt.Errorf("error scanning allocation type row: %v", err)
		}
		allocationTypeConfigs = append(allocationTypeConfigs, allocationTypeConfig)
	}

	// Check for any errors that occurred during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	// Return the list of allocation types
	return allocationTypeConfigs, nil
}
