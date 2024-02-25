package user

import (
	"context"
	"doremi/internal/model"
	"doremi/internal/repository/user"
)

type UserService struct {
	repo user.IUserRepository
}

type IUserService interface {
	CreateNewUser(user model.User, ctx context.Context) error
}

func NewUserService(repo user.IUserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) CreateNewUser(user model.User, ctx context.Context) error {
	return u.repo.CreateNewUser(ctx, user)
}
