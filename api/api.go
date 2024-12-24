package api

import (
	"fmt"
	"playlist-chat/spotify"

	"github.com/gorilla/mux"
)

const apiBasePath = "/api/v1"

type apiConfig struct {
	SpotifyClient *spotify.SpotifyClient
	Router        *mux.Router
	BasePath      string
}

func NewAPIConfig() *apiConfig {
	return &apiConfig{
		Router:   mux.NewRouter(),
		BasePath: apiBasePath,
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

func (a *apiConfig) PrintRoutes() error {
	err := a.Router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		var path string

		currentPath, err := route.GetPathTemplate()
		if err == nil {
			path = currentPath
		} else {
			path = "<unknown>"
		}

		methods, err := route.GetMethods()
		if err == mux.ErrMethodMismatch {
			methods = []string{"-ANY-"}
		} else if err != nil {
			methods = []string{"-NONE-"}
		}

		fmt.Printf("Path: %s, Methods: %v\n", path, methods)
		return nil
	})

	if err != nil {
		return fmt.Errorf("error printing routes: %w", err)
	}

	return nil
}
