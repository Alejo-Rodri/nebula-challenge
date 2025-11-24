package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	BaseApiURL string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		BaseApiURL: getEnv("BASE_API_URL", "https://api.ssllabs.com/api/v2"),
	}
}

// gets the env by key or fallbacks
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
