package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
)

var tmplt *template.Template
var ctx = context.Background()
var pg *Postgres

type News struct {
	Headline string
	Body     string
}

func main() {
	pg, _ = LoadDatabase(ctx)

	// // database.Ping(ctx)
	// pg.CreateUser(ctx, `a`, `b`, `c`)

	// result, _ := pg.GetUser(ctx, `a`)

	// fmt.Println(result)

	runServer()
}

func runServer() {
	http.HandleFunc("/home", handlePage)
	err := http.ListenAndServe("localhost:8080", nil)

	if err != nil {
		fmt.Errorf("There's an error with the server: %w", err)
	}
}

func handlePage(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		tmplt, _ = template.ParseFiles("main.html")

		result, _ := pg.GetUser(ctx, `a`)

		fmt.Println(result)

		event := News{
			Headline: "makeuseof.com has everything Tech",
			Body:     "Visit MUO for anything technology related",
		}

		fmt.Println(event)

		err := tmplt.Execute(writer, result)

		if err != nil {
			return
		}
	}
}
