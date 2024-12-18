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

	err = spotify.FetchClientToken(spotifyClient)
	if err != nil {
		return err
	}

	a.SpotifyClient = spotifyClient
	return nil
}
