package utils

import (
	"log"
	"os"
)

func GetEnvOrDie(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("failed to load env var: %s", key)
	}

	return val
}
