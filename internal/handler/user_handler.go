package handler

import (
	"context"
	"master-finanacial-planner/internal/entity"
	"master-finanacial-planner/internal/helper"
	"net/http"
)

func (h *Handler) SignUpHandler(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	response, err := h.userUsecases.SignUpUsecase(ctx, r)
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

func (h *Handler) SignInHandler(w http.ResponseWriter, r *http.Request) {

	// calling the usecase for the business logic
	ctx := context.Background()
	response, err := h.userUsecases.SignInUsecase(ctx, r)
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
