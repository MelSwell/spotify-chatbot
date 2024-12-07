package api

import "github.com/gorilla/mux"

func (a *apiConfig) RegisterRoutes() {
	r := a.Router.PathPrefix("/api/v1").Subrouter()

	a.registerSpotifyRoutes(r)
}

func (a *apiConfig) registerSpotifyRoutes(subrouter *mux.Router) {
	h := NewSpotifyHandler(a.SpotifyClient)
	r := subrouter.PathPrefix("/spotify").Subrouter()

	r.HandleFunc("/search", h.SearchSpotify).Methods("GET")
	r.HandleFunc("/search/playlists", h.SearchPlaylists).Methods("GET")
	r.HandleFunc("/search/tracks", h.SearchTracks).Methods("GET")
	r.HandleFunc("/search/albums", h.SearchAlbums).Methods("GET")
	r.HandleFunc("/search/artists", h.SearchArtists).Methods("GET")

}
