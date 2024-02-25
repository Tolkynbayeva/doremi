package service

import (
	"doremi/internal/repository"
	"doremi/internal/service/user"
)

type Service struct {
	User user.IUserService
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		User: user.NewUserService(repo),
	}
}
