package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	BaseApiURL string
	ConcurrentAssessments int
	ConcurrentRequestsRetry int
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		BaseApiURL: getEnv("BASE_API_URL", "https://api.ssllabs.com/api/v2"),
		ConcurrentAssessments: getEnvInt("CONCURRENT_ASSESSMENTS", 25),
		ConcurrentRequestsRetry: getEnvInt("CONCURRENT_REQUESTS_RETRY", 2),
	}
}

// gets the env by key or fallbacks
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
			intVal, err := strconv.Atoi(value)

		if err != nil {
			log.Fatalln("error parsing", key)
			return fallback
		}
		return intVal
	}

	return fallback
}
