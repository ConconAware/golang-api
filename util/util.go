package util

import (
	"fmt"

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

type error interface {
	Error() string
}

type errorString struct {
	Message string
}

func (e *errorString) Error() string {
	return e.Message
}

func New(text string) error {
	fmt.Println("Error :", text)
	return &errorString{text}
}
