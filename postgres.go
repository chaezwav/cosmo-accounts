package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/chaeyeonswav/cosmoAccounts/appconfig"
)

type Postgres struct {
	db *pgxpool.Pool
}

type User struct {
	wallet_address string
	nickname       string
	cosmo_id       string
}

var pgInstance *Postgres

func LoadDatabase(ctx context.Context) (*Postgres, error) {

	cfg, err := appconfig.LoadFromPath(context.Background(), "pkl/dev/config.pkl")
	if err != nil {
		panic(err)
	}

	cfPg := cfg.Postgres
	s := fmt.Sprintf(`postgresql://%v:%v@%v:%v/%v`, cfPg.Auth.Username, cfPg.Auth.Password, cfPg.Host, cfPg.Port, cfPg.Database)

	db, err := pgxpool.New(ctx, s)

	if err != nil {
		fmt.Errorf("couldn't connect to the database: %w", err)
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
	_, err := pg.db.Exec(ctx, `INSERT INTO dbo.user VALUES ($1, $2, $3)`, wa, n, cid)

	if err != nil {
		return fmt.Errorf("unable to create the user: %w", err)
	} else {
		fmt.Printf("Inserted user: %v\n", wa)
	}

	return nil
}

func (pg *Postgres) DeleteUser(ctx context.Context, u string) error {
	_, err := pg.db.Exec(ctx, `DELETE FROM dbo.user WHERE wallet_address = $1 RETURNING *;`, u)

	if err != nil {
		return fmt.Errorf("unable to delete the user: %w", err)
	} else {
		fmt.Printf("Deleted user: %v\n", u)
	}

	return nil
}

func (pg *Postgres) ModifyUser(ctx context.Context, u string, k string, nv string) error {
	_, err := pg.db.Exec(ctx, `UPDATE dbo.user SET $1 = $2 WHERE wallet_address = $3;`, k, nv, u)

	if err != nil {
		return fmt.Errorf("unable to modify the user: %w", err)
	} else {
		fmt.Printf("Modified user: %v\n", u)
	}

	return nil
}

func (pg *Postgres) GetUser(ctx context.Context, u string) (User, error) {
	var nick string
	var id string

	err := pg.db.QueryRow(ctx, `SELECT nickname, cosmo_id FROM dbo.user WHERE wallet_address = $1`, u).Scan(&nick, &id)
	if err != nil {
		fmt.Errorf("unable to get the user: %w", err)
	} else {
		fmt.Printf("Got user: %v\n", u)
	}

	return User{wallet_address: u, nickname: nick, cosmo_id: id}, nil
}

// TODO FETCH MANY user function
//* TODO More tables for foreign key relations
//* I.E: Applications
