package finance

import (
	"master-finanacial-planner/internal/repo"
)

type FinanceUsecase struct {
	financeRepo repo.ResourceRepo
}

func NewFinanceUsecase(dataResourceRepo repo.ResourceRepo) *FinanceUsecase {
	return &FinanceUsecase{
		financeRepo: dataResourceRepo,
	}
}
