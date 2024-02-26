package user

import (
	"context"
	"doremi/internal/model"
	"doremi/internal/pkg"
	"doremi/internal/repository/user"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type UserService struct {
	repo         user.IUserRepository
	jwtSecretKey string
}

type IUserService interface {
	CreateNewUser(user model.User, ctx context.Context) error
	CheckUserCredentials(creds model.User, ctx context.Context) (model.User, error)
	JwtAuth(user model.User) (string, error)
}

func NewUserService(repo user.IUserRepository, jwtSecret string) *UserService {
	return &UserService{
		repo:         repo,
		jwtSecretKey: jwtSecret,
	}
}

func (u *UserService) CreateNewUser(user model.User, ctx context.Context) error {
	var err error
	user.Password, err = pkg.HashPassword(user.Password)
	if err != nil {
		return err
	}

	return u.repo.CreateNewUser(ctx, user)
}

func (u *UserService) CheckUserCredentials(creds model.User, ctx context.Context) (model.User, error) {
	user, err := u.repo.GetUserByUsername(ctx, creds.Username)
	if err != nil {
		return model.User{}, err
	}

	if !pkg.CheckPasswordHash(creds.Password, user.Password) {
		return model.User{}, model.ErrIncorrectPassword
	}

	return user, nil
}

func (u *UserService) JwtAuth(user model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(u.jwtSecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
