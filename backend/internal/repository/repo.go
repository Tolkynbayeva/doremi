package repository

import (
	"doremi/internal/repository/user"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	user.IUserRepository
}

func NewRepo(client *mongo.Client) *Repository {
	return &Repository{
		IUserRepository: user.NewUserRepo(client),
	}
}
