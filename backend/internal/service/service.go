package service

import (
	"doremi/internal/repository"
	"doremi/internal/service/spotify"
	"doremi/internal/service/user"
)

type Service struct {
	User         user.IUserService
	Spotify      spotify.ISpotifyInterface
	JwtSecretKey string
}

func NewService(repo repository.Repository, jwtSecret string, spotifyID, spotifySecret string) *Service {
	return &Service{
		User:    user.NewUserService(repo, jwtSecret),
		Spotify: spotify.NewSpotifyService(spotifyID, spotifySecret),
	}
}
