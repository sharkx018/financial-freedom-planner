package handler

import (
	"context"
	"master-finanacial-planner/internal/entity"
	"net/http"
)

type FinanceUsecase interface {
	GetAssetClass(ctx context.Context, r *http.Request) (*entity.ApiResponse, error)
	GetEffectiveReturnAllocationType(ctx context.Context, r *http.Request) (*entity.ApiResponse, error)
	GetInvestingSurplus(ctx context.Context, r *http.Request) (*entity.ApiResponse, error)
	GetNetWorth(ctx context.Context, r *http.Request) (*entity.ApiResponse, error)
	//GetSipAllocation(ctx context.Context, r *http.Request) (*entity.ApiResponse, error)
}

type UserUsecases interface {
	SignUpUsecase(ctx context.Context, r *http.Request) (*entity.ApiResponse, error)
	SignInUsecase(ctx context.Context, r *http.Request) (*entity.ApiResponse, error)
}

type Handler struct {
	financeUsecases FinanceUsecase
	userUsecases    UserUsecases
}

func NewFinanceHandler(userUsecases UserUsecases, financeUsecases FinanceUsecase) *Handler {
	return &Handler{
		userUsecases:    userUsecases,
		financeUsecases: financeUsecases,
	}
}
