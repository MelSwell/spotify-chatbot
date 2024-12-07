package api

import (
	"playlist-chat/spotify"

	"github.com/gorilla/mux"
)

type apiConfig struct {
	SpotifyClient *spotify.SpotifyClient
	Router        *mux.Router
}

func NewAPIConfig() *apiConfig {
	return &apiConfig{
		Router: mux.NewRouter(),
	}
}

func (a *apiConfig) InitSpotifyClient(spotifyClientID, spotifyClientSecret string) error {
	spotifyClient, err := spotify.NewSpotifyClient(spotifyClientID, spotifyClientSecret)
	if err != nil {
		return err
	}
	tok, err := spotifyClient.GetClientToken()
	if err != nil {
		return err
	}
	spotifyClient.ClientAccessToken = tok
	a.SpotifyClient = spotifyClient
	return nil
}
