package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"

	"github.com/chaeyeonswav/cosmoAccounts/appconfig"
)

type Postgres struct {
	db *pgx.Conn
}

var pgInstance *Postgres

func LoadDatabase(ctx context.Context) (*Postgres, error) {

	cfg, err := appconfig.LoadFromPath(context.Background(), "pkl/dev/config.pkl")
	if err != nil {
		panic(err)
	}

	cfPg := cfg.Postgres
	s := fmt.Sprintf(`postgresql://%v:%v@%v:%v/%v`, cfPg.Auth.Username, cfPg.Auth.Password, cfPg.Host, cfPg.Port, cfPg.Database)

	db, err := pgx.Connect(ctx, s)
	if err != nil {
		fmt.Errorf("couldn't connect to the database: %w", err)
		os.Exit(1)
	}

	query := `
	CREATE SCHEMA IF NOT EXISTS dbo AUTHORIZATION @user;
	CREATE TABLE IF NOT EXISTS dbo.user (
		wallet_address        text UNIQUE PRIMARY KEY,
		nickname         	  text UNIQUE,
		cosmo_id        	  text UNIQUE
	);
	`

	_, qerr := db.Exec(ctx, query, pgx.NamedArgs{"user": cfPg.Auth.Username})
	if err != nil {
		fmt.Errorf("couldn't create the database table: %w", qerr)
	}

	pgInstance = &Postgres{db}
	return pgInstance, nil
}

func (pg *Postgres) Ping(ctx context.Context) error {
	err := pg.db.Ping(ctx)

	if err != nil {
		return fmt.Errorf("unable to ping the database: %w", err)
	} else {
		fmt.Println("Pinged database")
	}

	return nil
}

func (pg *Postgres) CreateUser(ctx context.Context, wa string, n string, cid string) error {
	result, err := pg.db.Exec(ctx, `INSERT INTO dbo.user VALUES ($1, $2, $3)`, wa, n, cid)

	if err != nil {
		return fmt.Errorf("unable to insert create user: %w", err)
	} else {
		fmt.Printf("Inserted user: %v\n", result)
	}

	return nil
}

// TODO DELETE, MODIFY, FETCH user function
//* TODO More tables for foreign key relations
//* I.E: Applications
