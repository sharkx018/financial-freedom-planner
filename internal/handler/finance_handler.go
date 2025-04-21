package handler

import (
	"context"
	"master-finanacial-planner/internal/entity"
	"master-finanacial-planner/internal/helper"
	"net/http"
)

func (h *Handler) GetAssetClassHandler(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	response, err := h.financeUsecases.GetAssetClass(ctx, r)
	if err != nil {
		rr := &entity.ApiResponse{
			Data: nil,
			Error: &entity.CommonErrorResponse{
				Message: err.Error(),
			},
		}
		helper.WriteCustomResp(w, 500, rr)
	} else {
		response.Error = nil
		response.Success = true
		helper.WriteCustomResp(w, http.StatusOK, response)
	}

}

func (h *Handler) GetEffectiveReturnAllocationTypeHandler(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	response, err := h.financeUsecases.GetEffectiveReturnAllocationType(ctx, r)
	if err != nil {
		rr := &entity.ApiResponse{
			Data: nil,
			Error: &entity.CommonErrorResponse{
				Message: err.Error(),
			},
		}
		helper.WriteCustomResp(w, 500, rr)
	} else {
		response.Error = nil
		response.Success = true
		helper.WriteCustomResp(w, http.StatusOK, response)
	}

}

func (h *Handler) GetInvestingSurplusHandler(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	response, err := h.financeUsecases.GetInvestingSurplus(ctx, r)
	if err != nil {
		rr := &entity.ApiResponse{
			Data: nil,
			Error: &entity.CommonErrorResponse{
				Message: err.Error(),
			},
		}
		helper.WriteCustomResp(w, 500, rr)
	} else {
		response.Error = nil
		response.Success = true
		helper.WriteCustomResp(w, http.StatusOK, response)
	}

}

func (h *Handler) GetNetWorthHandler(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	response, err := h.financeUsecases.GetNetWorth(ctx, r)
	if err != nil {
		rr := &entity.ApiResponse{
			Data: nil,
			Error: &entity.CommonErrorResponse{
				Message: err.Error(),
			},
		}
		helper.WriteCustomResp(w, 500, rr)
	} else {
		response.Error = nil
		response.Success = true
		helper.WriteCustomResp(w, http.StatusOK, response)
	}

}

func (h *Handler) SipAllocatorHandler(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	response, err := h.financeUsecases.SipAllocator(ctx, r)
	if err != nil {
		rr := &entity.ApiResponse{
			Data: nil,
			Error: &entity.CommonErrorResponse{
				Message: err.Error(),
			},
		}
		helper.WriteCustomResp(w, 500, rr)
	} else {
		response.Error = nil
		response.Success = true
		helper.WriteCustomResp(w, http.StatusOK, response)
	}

}

func (h *Handler) GetInvestableAssetAllocation(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	response, err := h.financeUsecases.GetInvestableAssetAllocation(ctx, r)
	if err != nil {
		rr := &entity.ApiResponse{
			Data: nil,
			Error: &entity.CommonErrorResponse{
				Message: err.Error(),
			},
		}
		helper.WriteCustomResp(w, 500, rr)
	} else {
		response.Error = nil
		response.Success = true
		helper.WriteCustomResp(w, http.StatusOK, response)
	}

}
