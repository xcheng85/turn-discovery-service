package utils

import (
	"log"
	"os"
)

func GetEnvVar(name string, isRequired bool, defaultValue string) string {
	value := os.Getenv(name)
	if isRequired && len(value) == 0 {
		log.Fatalf("Missing env var: %s", name)
	} else if len(value) == 0 {
		value = defaultValue
	}
	return value
}
