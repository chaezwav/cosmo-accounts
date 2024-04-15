package main

import (
	"context"
)

func main() {
	ctx := context.Background()

	pg, _ := LoadDatabase(ctx)

	// database.Ping(ctx)
	pg.CreateUser(ctx, `a`, `b`, `c`)
}
