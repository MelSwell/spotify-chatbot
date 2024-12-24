package spotify

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	BaseURL            = "https://api.spotify.com/v1"
	SearchTypeAlbum    = "album"
	SearchTypeArtist   = "artist"
	SearchTypePlaylist = "playlist"
	SearchTypeTrack    = "track"
	SearchTypeAll      = "album,artist,playlist,track"
)

type SpotifyClient struct {
	ClientToken  *SpotifyToken
	ClientID     string
	ClientSecret string
	Credentials  string
}

func NewSpotifyClient(clientID, clientSecret string) (*SpotifyClient, error) {
	return &SpotifyClient{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Credentials:  clientID + ":" + clientSecret,
	}, nil
}

func (s *SpotifyClient) Search(query string, typeParam string) (result CombinedResponse, err error) {
	q := url.QueryEscape(query)
	url := fmt.Sprintf("%s/search?type=%s&q=%s", BaseURL, typeParam, q)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+s.ClientToken.AccessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return CombinedResponse{}, err
	}
	defer resp.Body.Close()

	// use this until we figure out how to handle spotify errors
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return CombinedResponse{}, fmt.Errorf("search failed, status: %s, body: %s", resp.Status, body)
	}

	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return CombinedResponse{}, err
	}

	return result, nil
}

func (s *SpotifyClient) Test() (Playlist, error) {
	req, _ := http.NewRequest("GET", "https://api.spotify.com/v1/playlists/03irz9t12mNseVaUUzb1UA", nil)
	req.Header.Set("Authorization", "Bearer "+s.ClientToken.AccessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Playlist{}, err
	}
	defer resp.Body.Close()

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return nil, err
	// }

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return Playlist{}, fmt.Errorf("search failed, status: %s, body: %s", resp.Status, body)
	}

	var body Playlist
	if err = json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return Playlist{}, err
	}
	return body, nil
}

func (s *SpotifyClient) SearchPlaylists(query string) (SearchResponse[SearchItemPlaylist], error) {
	result, err := s.Search(query, SearchTypePlaylist)
	if err != nil {
		return SearchResponse[SearchItemPlaylist]{}, err
	}

	return result.Playlists, nil
}

func (s *SpotifyClient) SearchTracks(query string) (SearchResponse[SearchItemTrack], error) {
	result, err := s.Search(query, SearchTypeTrack)
	if err != nil {
		return SearchResponse[SearchItemTrack]{}, err
	}

	return result.Tracks, nil
}

func (s *SpotifyClient) SearchArtists(query string) (SearchResponse[SearchItemArtist], error) {
	result, err := s.Search(query, SearchTypeArtist)
	if err != nil {
		return SearchResponse[SearchItemArtist]{}, err
	}

	return result.Artists, nil
}

func (s *SpotifyClient) SearchAlbums(query string) (SearchResponse[SearchItemAlbum], error) {
	result, err := s.Search(query, SearchTypeAlbum)
	if err != nil {
		return SearchResponse[SearchItemAlbum]{}, err
	}

	return result.Albums, nil
}
