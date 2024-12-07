package spotify

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type SpotifyClient struct {
	ClientAccessToken string
	ClientID          string
	ClientSecret      string
	Credentials       string
}

func NewSpotifyClient(clientID, clientSecret string) (*SpotifyClient, error) {
	return &SpotifyClient{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Credentials:  clientID + ":" + clientSecret,
	}, nil
}

func (s *SpotifyClient) GetClientToken() (token string, err error) {
	authURL := "https://accounts.spotify.com/api/token"
	data := []byte("grant_type=client_credentials")
	req, err := http.NewRequest("POST", authURL, bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(s.Credentials)))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get spotify client access token, status: %s", resp.Status)
	}

	var body struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return "", err
	}

	return body.AccessToken, nil
}

func (s *SpotifyClient) Search(query string, typeParam string) ([]byte, error) {
	encodedQ := url.QueryEscape(query)
	var url string
	if typeParam == "" {
		url = "https://api.spotify.com/v1/search?type=album,artist,playlist,track&q=" + encodedQ
	} else {
		url = "https://api.spotify.com/v1/search?type=" + typeParam + "&q=" + encodedQ
	}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+s.ClientAccessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search failed, status: %s, body: %s", resp.Status, body)
	}

	return body, nil
}

func (s *SpotifyClient) SearchPlaylists(query string) ([]byte, error) {
	return s.Search(query, "playlist")
}

func (s *SpotifyClient) SearchTracks(query string) ([]byte, error) {
	return s.Search(query, "track")
}

func (s *SpotifyClient) SearchArtists(query string) ([]byte, error) {
	return s.Search(query, "artist")
}

func (s *SpotifyClient) SearchAlbums(query string) ([]byte, error) {
	return s.Search(query, "album")
}
