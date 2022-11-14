package util

import (
	"github.com/joho/godotenv"

	"log"
	"os"
)

func LoadEnv() {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
}

func GetEnv(name string) string {
	return os.Getenv(name)
}
