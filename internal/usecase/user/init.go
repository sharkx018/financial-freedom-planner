package user

import (
	"context"
	"master-finanacial-planner/internal/entity"
	"master-finanacial-planner/internal/repo"
	"net/http"
)

type UserUsecase struct {
	userRepo repo.ResourceRepo
}

func (u UserUsecase) SignUpUsecase(ctx context.Context, r *http.Request) (*entity.ApiResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserUsecase) SignInUsecase(ctx context.Context, r *http.Request) (*entity.ApiResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserUsecase(dataResourceRepo repo.ResourceRepo) *UserUsecase {
	return &UserUsecase{
		userRepo: dataResourceRepo,
	}
}
