package service

import (
	"context"
	"webook/internal/domain"
	"webook/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo}
}

func (svc *UserService) SignUp(ctx context.Context, u domain.User) error {
	// 考虑加密
	// 然后存起来
	return svc.repo.Create(ctx, u)
}
