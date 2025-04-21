package finance

import (
	"context"
	"master-finanacial-planner/internal/entity"
	"master-finanacial-planner/internal/helper"
	"net/http"
)

func (f FinanceUsecase) GetEffectiveReturnAllocationType(ctx context.Context, r *http.Request) (*entity.ApiResponse, error) {

	result, err := f.getAllocationTypeReturns(ctx)
	if err != nil {
		return nil, err
	}

	return &entity.ApiResponse{
		Data: map[string]interface{}{
			"message":           "Asset class data fetched successfully",
			"effective-returns": result,
		},
		Success: true,
	}, nil

}

func (f FinanceUsecase) getAllocationTypeReturns(ctx context.Context) (map[string]float64, error) {
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
		result[k] = helper.RoundToDecimals(v, 1)
	}

	return result, nil
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

func (f FinanceUsecase) SipAllocator(ctx context.Context, r *http.Request) (*entity.ApiResponse, error) {

	// get the goal
	goalsData, err := f.financeRepo.GetGoals(ctx)
	if err != nil {
		return nil, err
	}

	var sipAllocator = make(map[string]float64)

	// for each goal
	for _, goal := range goalsData {

		// calculate the inflated amount
		inflatedAmount := helper.InflationCalculator(goal.TodayAmount, goal.YearsLeft, goal.InflationPercentage)
		// subtract the allocated amount
		requiredAmount := inflatedAmount - goal.AllocatedAmount

		// calculate the sip required
		if requiredAmount <= 0 {
			continue
		}

		allocationTypeReturnsMap, err := f.getAllocationTypeReturns(ctx)
		if err != nil {
			return nil, err
		}

		sipRequired := helper.CalculateSIPRequired(requiredAmount, goal.YearsLeft, allocationTypeReturnsMap[goal.Name], goal.SIPStepUpPercentage)

		// get the allocation type from the years left
		allocationData, err := f.financeRepo.GetAllocationByYearLeft(ctx, goal.YearsLeft)
		if err != nil {
			return nil, err
		}

		// get allocation type config
		allocationConfigData, err := f.financeRepo.GetAllocationConfigByAllocationTypeId(ctx, allocationData[0].ID)
		if err != nil {
			return nil, err
		}

		// divide the sip amount according to the asset class
		for _, assetAllocationInfo := range allocationConfigData {
			sipAllocator[assetAllocationInfo.AssetName] += sipRequired * assetAllocationInfo.AllocationInPercentage / 100
		}

	}

	return &entity.ApiResponse{
		Data: map[string]interface{}{
			"message":       "SIP allocation fetched successfully",
			"Sip Allocator": sipAllocator,
		},
		Success: true,
	}, nil

}
