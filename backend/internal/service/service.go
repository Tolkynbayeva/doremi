package service

import (
	"doremi/internal/repository"
	"doremi/internal/service/user"
)

type Service struct {
	User         user.IUserService
	JwtSecretKey string
}

func NewService(repo repository.Repository, jwtSecret string) *Service {
	return &Service{
		User: user.NewUserService(repo, jwtSecret),
	}
}
