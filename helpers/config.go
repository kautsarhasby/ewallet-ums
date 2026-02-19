package helpers

import (
	"os"

	"github.com/joho/godotenv"
)

var Env map[string]string

func SetupEnv() {
	_ = godotenv.Load(".env")
}

func GetEnv(key, value string) string {
	val := os.Getenv(key)
	if val == "" {
		val = value
	}
	return val
}
