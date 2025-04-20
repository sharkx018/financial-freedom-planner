package repo

import (
	"context"
	"database/sql"
	"fmt"
	"master-finanacial-planner/internal/logger"
)

type AssetClass struct {
	ID                         int64   `json:"id"`                            // bigint corresponds to int64 in Go
	Name                       string  `json:"name"`                          // varchar corresponds to string
	ExpectedReturnInPercentage float64 `json:"expected_return_in_percentage"` // double precision corresponds to float64
}

type ResourceRepo interface {
	GetAssetClass(ctx context.Context) ([]AssetClass, error)
}

type ResourceRepository struct {
	db *sql.DB
}

func NewResource(db *sql.DB) *ResourceRepository {
	return &ResourceRepository{
		db,
	}
}

func (r *ResourceRepository) GetAssetClass(ctx context.Context) ([]AssetClass, error) {
	var assetClasses []AssetClass

	query := "SELECT id, name, expected_return_in_percentage FROM asset_class"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error querying asset class data: %v", err)
	}
	defer rows.Close() // Ensure the rows iterator is closed properly

	for rows.Next() {
		var assetClass AssetClass
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
