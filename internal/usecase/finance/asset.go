package finance

import (
	"context"
	"master-finanacial-planner/internal/entity"
	"net/http"
)

func (f FinanceUsecase) GetAssetClass(ctx context.Context, r *http.Request) (*entity.ApiResponse, error) {
	data, err := f.financeRepo.GetAssetClass(ctx)
	if err != nil {
		return nil, err
	}

	return &entity.ApiResponse{
		Data: map[string]interface{}{
			"message": "Asset class data fetched successfully",
			"data":    data,
		},
		Success: true,
	}, nil
}
