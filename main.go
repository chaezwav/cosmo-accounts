package main

import (
	"context"
)

func main() {
	ctx := context.Background()

	database, _ := LoadDatabase(ctx)

	database.CreateUser(ctx, `hi`, `hi`, `hi`)
}
