package main

import (
	"log"
	"net/http"
	"os"
	"playlist-chat/api"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	spotifyClientID := os.Getenv("SPOTIFY_CLIENT_ID")
	spotifyClientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")
	if spotifyClientID == "" || spotifyClientSecret == "" {
		log.Fatal("SPOTIFY_CLIENT_ID and SPOTIFY_CLIENT_SECRET environment variables are required")
	}

	api := api.NewAPIConfig()

	if err := api.InitSpotifyClient(spotifyClientID, spotifyClientSecret); err != nil {
		log.Fatal(err)
	}
	api.RegisterRoutes()

	log.Println("Server started on: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", api.Router))
}
