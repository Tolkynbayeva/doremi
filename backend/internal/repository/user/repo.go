package user

import (
	"context"
	"doremi/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	db *mongo.Database
}

type IUserRepository interface {
	CreateNewUser(ctx context.Context, user model.User) error
	DeleteUserById(ctx context.Context, id int) error
	GetUserById(ctx context.Context, id int) (model.User, error)
}

func NewUserRepo(db *mongo.Database) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) CreateNewUser(ctx context.Context, user model.User) error {
	collection := u.db.Collection("user")
	_, err := collection.InsertOne(ctx, user)
	return err
}

func (u *UserRepository) DeleteUserById(ctx context.Context, id int) error {
	collection := u.db.Collection("user")
	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (u *UserRepository) GetUserById(ctx context.Context, id int) (model.User, error) {
	var user model.User
	collection := u.db.Collection("user")
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	return user, err
}
