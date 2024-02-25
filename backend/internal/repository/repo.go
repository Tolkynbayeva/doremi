package repository

import (
	"doremi/internal/repository/user"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	user.IUserRepository
}

func NewRepo(client *mongo.Database) *Repository {
	return &Repository{
		IUserRepository: user.NewUserRepo(client),
	}
}
