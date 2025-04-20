package finance

import (
	"context"
	"master-finanacial-planner/internal/entity"
	"math"
	"net/http"
)

func (f FinanceUsecase) GetEffectiveReturnAllocationType(ctx context.Context, r *http.Request) (*entity.ApiResponse, error) {
	data, err := f.financeRepo.GetAllAllocationTypeConfig(ctx)
	if err != nil {
		return nil, err
	}

	result := make(map[string]float64)

	for _, row := range data {
		x := row.AssetReturns * row.AllocationInPercentage / 100
		result[row.AllocationTypeName] += x
	}

	result["medium-term"] = result["short-term"]*0.6 + result["medium-term"]*0.4

	for k, v := range result {
		result[k] = roundToTwoDecimals(v)
	}

	return &entity.ApiResponse{
		Data: map[string]interface{}{
			"message":           "Asset class data fetched successfully",
			"effective-returns": result,
		},
		Success: true,
	}, nil

}

func (f FinanceUsecase) GetInvestingSurplus(ctx context.Context, r *http.Request) (*entity.ApiResponse, error) {
	data, err := f.financeRepo.GetInvestingSurplus(ctx)
	if err != nil {
		return nil, err
	}

	return &entity.ApiResponse{
		Data: map[string]interface{}{
			"message":           "Investing surplus data fetched successfully",
			"investing-surplus": data,
		},
		Success: true,
	}, nil

}

func (f FinanceUsecase) GetNetWorth(ctx context.Context, r *http.Request) (*entity.ApiResponse, error) {

	// liquid and Illiquid
	data, err := f.financeRepo.GetLiquidAndIlliquidAssets(ctx)
	if err != nil {
		return nil, err
	}

	// liabilities
	liabilitiesAmount, err := f.financeRepo.GetAllLiability(ctx)
	if err != nil {
		return nil, err
	}

	var totalAsset float64
	var liquidAsset float64

	for k, v := range data {
		totalAsset += v

		if k == "liquid" {
			liquidAsset += v
		}
	}

	netWorth := totalAsset - liabilitiesAmount

	return &entity.ApiResponse{
		Data: map[string]interface{}{
			"message":      "Net Worth info fetched successfully",
			"total_asset":  totalAsset,
			"liquid_asset": liquidAsset,
			"net_worth":    netWorth,
		},
		Success: true,
	}, nil

}

func roundToTwoDecimals(value float64) float64 {
	return math.Round(value*10) / 10
}

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
