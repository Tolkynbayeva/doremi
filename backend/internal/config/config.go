package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Host    string
	Port    string
	Auth    Auth
	MongoDB MongoDB
}

type MongoDB struct {
	User     string
	Password string
	Host     string
	Name     string
}

type Auth struct {
	JwtSecretKey string
}

func NewConfig(path string) (Config, error) {
	if err := godotenv.Load(path); err != nil {
		return Config{}, err
	}

	cfg := Config{
		Host: os.Getenv("HOST"),
		Port: os.Getenv("PORT"),
		MongoDB: MongoDB{
			User:     os.Getenv("MONGO_USER"),
			Password: os.Getenv("MONGO_PSW"),
			Host:     os.Getenv("MONGO_HOST"),
			Name:     os.Getenv("MONGO_DB_NAME"),
		},
		Auth: Auth{
			JwtSecretKey: os.Getenv("JWT_SECRET_KEY"),
		},
	}

	return cfg, nil
}
