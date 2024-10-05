package config

import (
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	DbPath         string
	YouTubeAPIKeys []string
	SearchQuery    string
	FetchInterval  time.Duration
}

// Returns the env value if found else uses a default value
func getEnv(key, def string) string {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	return val
}

// returns a config struct
func Load() *Config {
	godotenv.Load()
	return &Config{
		DbPath:         getEnv("DB_PATH", "test.db"),
		YouTubeAPIKeys: strings.Split(os.Getenv("YOUTUBE_API_KEYS"), ","),
		SearchQuery:    getEnv("SEARCH_QUERY", "marvel"),
		FetchInterval:  2 * time.Hour,
	}
}
