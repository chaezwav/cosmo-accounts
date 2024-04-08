package main

import (
	"context"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	pgInstance, _ = LoadDatabase(context.Background(), os.Getenv("DATABASE_URL"))

	pgInstance.Ping(context.Background())
	pgInstance.CreateUser(context.Background(), "0xFD1093D8FC8d8596882549dC3aeC611a8F700390", "koehn", "chaezwav")
}
