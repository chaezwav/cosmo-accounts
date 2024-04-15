package main

import (
	"context"
	"fmt"
	"os"

	"github.com/edgedb/edgedb-go"
)

type EdgeDB struct {
	db *edgedb.Client
}

var (
	edbInstance *EdgeDB
)

func LoadDatabase(ctx context.Context) (*EdgeDB, error) {
	client, err := edgedb.CreateClient(ctx, edgedb.Options{})
	if err != nil {
		fmt.Errorf("couldn't connect to the database: %w", err)
		os.Exit(1)
	}

	edbInstance = &EdgeDB{db: client}
	return edbInstance, err
}
