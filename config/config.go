package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

const (
	SpotifyTokenURL    = "https://accounts.spotify.com/api/token"
	SpotifyPlaylistURL = "https://api.spotify.com/v1/playlists/%s/tracks"
)

var HibanaConcert = time.Date(2025, time.May, 8, 0, 0, 0, 0, time.UTC)

var AppConfig Config

type Config struct {
	Token               string
	AppID               string
	SpotifyClientID     string
	SpotifyClientSecret string
	SpotifyPlaylistID   string
}

func LoadConfig() Config {
	godotenv.Load()

	return Config{
		Token:               getEnv("DISCORD_TOKEN", ""),
		AppID:               getEnv("DISCORD_APP_ID", ""),
		SpotifyClientID:     getEnv("SPOTIFY_CLIENT_ID", ""),
		SpotifyClientSecret: getEnv("SPOTIFY_CLIENT_SECRET", ""),
		SpotifyPlaylistID:   getEnv("SPOTIFY_PLAYLIST_ID", "5J0AuCTRCc71qE7klIbjmA"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func init() {
	AppConfig = LoadConfig()
}
