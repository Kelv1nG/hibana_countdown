package spotify

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"

	"github.com/Kelv1nG/hibana_countdown/config"
)

func GetSpotifyToken() (string, error) {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", config.SpotifyTokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(config.AppConfig.SpotifyClientID, config.AppConfig.SpotifyClientSecret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	token, ok := result["access_token"].(string)
	if !ok {
		return "", fmt.Errorf("failed to get access token")
	}

	return token, nil
}

func GetSongLinks(playlistID, spotifyToken string) ([]string, error) {
	url := fmt.Sprintf(config.SpotifyPlaylistURL, playlistID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+spotifyToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	items, ok := result["items"].([]any)
	if !ok {
		return nil, fmt.Errorf("failed to parse tracks, 'items' field missing or not an array")
	}

	trackURLs := make([]string, 0, len(items))
	for _, item := range items {
		trackData, ok := item.(map[string]any)
		if !ok {
			continue
		}
		track, ok := trackData["track"].(map[string]any)
		if !ok {
			continue
		}
		externalURLs, ok := track["external_urls"].(map[string]any)
		if !ok {
			continue
		}
		spotifyURL, ok := externalURLs["spotify"].(string)
		if !ok {
			continue
		}
		trackURLs = append(trackURLs, spotifyURL)
	}

	return trackURLs, nil
}

func RandomHibanaSong() (string, error) {
	token, err := GetSpotifyToken()
	if err != nil {
		log.Printf("Failed to get Spotify token: %v", err)
		return "", err
	}

	trackURLs, err := GetSongLinks(config.AppConfig.SpotifyPlaylistID, token)
	if err != nil {
		log.Printf("Failed to get song links: %v", err)
		return "", err
	}

	randTrackURL := trackURLs[rand.Intn(len(trackURLs))]
	return randTrackURL, nil
}
