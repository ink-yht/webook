package service

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"webook/internal/domain"
	"webook/internal/repository"
)

var ErrUserDuplicateEmail = repository.ErrUserDuplicateEmail

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo}
}

func (svc *UserService) SignUp(ctx context.Context, u domain.User) error {
	// 考虑加密
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	// 然后存起来
	return svc.repo.Create(ctx, u)
}
