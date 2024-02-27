package spotify

type SpotifyService struct {
	clientID  string
	secretKey string
}

type ISpotifyInterface interface {
	GetTopTracks()
}

func NewSpotifyService(clientID, key string) *SpotifyService {
	return &SpotifyService{
		clientID:  clientID,
		secretKey: key,
	}
}

func (s *SpotifyService) GetTopTracks() {

}
