package user

import (
	"context"
	"doremi/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	client *mongo.Client
}

type IUserRepository interface {
	CreateNewUser(ctx context.Context, user model.User) error
	DeleteUserById(ctx context.Context, id int) error
	GetUserById(ctx context.Context) (model.User, error)
}

func NewUserRepo(client *mongo.Client) *UserRepository {
	return &UserRepository{
		client: client,
	}
}

func (u *UserRepository) CreateNewUser(ctx context.Context, user model.User) error {
	return nil
}

func (u *UserRepository) DeleteUserById(ctx context.Context, id int) error {
	return nil
}

func (u *UserRepository) GetUserById(ctx context.Context) (model.User, error) {
	return model.User{}, nil
}
