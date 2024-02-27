package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Host    string
	Port    string
	Key     Key
	MongoDB MongoDB
}

type MongoDB struct {
	User     string
	Password string
	Host     string
	Name     string
}

type Key struct {
	JwtSecretKey     string
	SpotifyID        string
	SpotifySecretKey string
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
		Key: Key{
			JwtSecretKey:     os.Getenv("JWT_SECRET_KEY"),
			SpotifyID:        os.Getenv("SPOTIFY_CLIENT_ID"),
			SpotifySecretKey: os.Getenv("SPOTIFY_SECRET_KEY"),
		},
	}

	return cfg, nil
}
