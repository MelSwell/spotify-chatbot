package api

import (
	"errors"
	"net/http"
	"playlist-chat/spotify"
)

type SpotifyHandler struct {
	Client *spotify.SpotifyClient
}

func NewSpotifyHandler(client *spotify.SpotifyClient) *SpotifyHandler {
	return &SpotifyHandler{Client: client}
}

func (s *SpotifyHandler) SearchSpotify(w http.ResponseWriter, r *http.Request) {
	q, err := extractQuery(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	searchResults, err := s.Client.Search(q, "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(searchResults))
}

func (s *SpotifyHandler) SearchPlaylists(w http.ResponseWriter, r *http.Request) {
	q, err := extractQuery(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	searchResults, err := s.Client.SearchPlaylists(q)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(searchResults))
}

func (s *SpotifyHandler) SearchTracks(w http.ResponseWriter, r *http.Request) {
	q, err := extractQuery(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	searchResults, err := s.Client.SearchTracks(q)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(searchResults))
}

func (s *SpotifyHandler) SearchAlbums(w http.ResponseWriter, r *http.Request) {
	q, err := extractQuery(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	searchResults, err := s.Client.SearchAlbums(q)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(searchResults))
}

func (s *SpotifyHandler) SearchArtists(w http.ResponseWriter, r *http.Request) {
	q, err := extractQuery(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	searchResults, err := s.Client.SearchArtists(q)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(searchResults))
}

func extractQuery(r *http.Request) (string, error) {
	query := r.URL.Query().Get("q")
	if query == "" {
		return "", errors.New("query parameter is required")
	}
	return query, nil
}
