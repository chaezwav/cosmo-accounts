package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	pg, _ := LoadDatabase(ctx)

	// database.Ping(ctx)
	pg.CreateUser(ctx, `a`, `b`, `c`)

	result, _ := pg.GetUser(ctx, `a`)

	fmt.Println(result)
}
