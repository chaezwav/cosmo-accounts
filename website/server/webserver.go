package server

// FIXME Compiler returns memory error likely due to templating

import (
	"context"
	"fmt"
	"html/template"
	"net/http"

	postgres "github.com/chaeyeonswav/cosmo-accounts/internal/database"
	"github.com/chaeyeonswav/cosmo-accounts/tools/pkl"
)

var tmplt *template.Template
var pg *postgres.Postgres

func RunServer() error {
	pg, _ = postgres.LoadDatabase(context.Background())

	pklcf, _ := pkl.LoadFromPath(context.Background(), "config/pkl/dev/config.pkl")
	http.HandleFunc("/home", handlePage)

	l := fmt.Sprintf("%v:%v", pklcf.Host, pklcf.Port)

	fmt.Println("Server is running on", l)

	err := http.ListenAndServe(l, nil)

	if err != nil {
		fmt.Errorf("There's an error with the server: %w", err)
	}

	fmt.Printf("Server is running on %v\n", l)

	return nil
}

func handlePage(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		tmplt, _ = template.ParseFiles("../tmpl/main.html")

		result, _ := pg.GetUser(context.Background(), `a`)

		err := tmplt.Execute(writer, result)

		if err != nil {
			return
		}
	}
}
