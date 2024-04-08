package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"sync"
)

type Postgres struct {
	db *pgxpool.Pool
}

var (
	pgInstance *Postgres
	pgOnce     sync.Once
)

func LoadDatabase(ctx context.Context, connString string) (*Postgres, error) {
	pgOnce.Do(func() {
		db, err := pgxpool.New(ctx, connString)
		if err != nil {
			fmt.Errorf("unable to create connection pool: %w", err)
		} else {
			fmt.Println("Successfully connected to database!")
		}

		pgInstance = &Postgres{db}
	})

	return pgInstance, nil
}

func (pg *Postgres) Ping(ctx context.Context) error {
	err := pg.db.Ping(ctx)

	if err != nil {
		return fmt.Errorf("error pinging database: %w\n", err)
	} else {
		fmt.Println("Pinged successfully!")
	}

	return err
}

func (pg *Postgres) Close() {
	pg.db.Close()
}

func (pg *Postgres) CreateUser(ctx context.Context, walletAddress string, nickname string, cosmoID string) error {

	_, err := pg.db.Exec(ctx, `INSERT INTO dbo.users (wallet_address, nickname, cosmo_id) VALUES (@address, @nickname, @id)`, pgx.NamedArgs{
		"address":  walletAddress,
		"nickname": nickname,
		"id":       cosmoID,
	})
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}
