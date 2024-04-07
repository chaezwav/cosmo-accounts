package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var walletAddress, nickname string

	err = conn.QueryRow(context.Background(), "SELECT wallet_address, nickname FROM dbo.users;").Scan(&walletAddress, &nickname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to insert into the table: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Address: %v, Nick: %v\n", walletAddress, nickname)
}
