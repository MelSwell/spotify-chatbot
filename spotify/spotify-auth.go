package spotify

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
)

type SpotifyToken struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

func FetchClientToken(s *SpotifyClient) (err error) {
	authURL := "https://accounts.spotify.com/api/token"
	data := []byte("grant_type=client_credentials")
	req, err := http.NewRequest("POST", authURL, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	creds := base64.StdEncoding.EncodeToString([]byte(s.Credentials))
	req.Header.Set("Authorization", "Basic "+creds)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to get spotify client access token, status: %s", resp.Status)
	}

	var token SpotifyToken
	if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {
		return err
	}

	s.ClientToken = &token
	return nil
}
