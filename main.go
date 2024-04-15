package main

import (
	"context"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Errorf("failed to load .env files: %w", err)
	}

	ctx := context.Background()

	database, err := LoadDatabase(ctx)

	var result string
	err = database.db.Execute(ctx, `
	INSERT User {
		wallet_address := 'hi'
	}
	`, &result)
	if err != nil {
		fmt.Errorf("failed to execute: %w", err)
	}

	fmt.Println(result)
}
