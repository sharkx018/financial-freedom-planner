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
	//Get(ctx context.Context) ([]AssetClass, error)
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
