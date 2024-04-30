package main

import (
	"design-pattern/configuration"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

const port = "localhost:4000"

type application struct {
	templateMap map[string]*template.Template
	config      appConfig
	App         *configuration.Application
}

type appConfig struct {
	useCache bool
	dsn      string
}

func main() {
	app := application{
		templateMap: make(map[string]*template.Template),
	}

	flag.BoolVar(&app.config.useCache, "cache", false, "Use template cache.")
	flag.StringVar(&app.config.dsn, "dsn", "mariadb:myverysecretpassword@tcp(localhost:3306)/breeders?parseTime=true&tls=false&collation=utf8_unicode_ci&timeout=5s", "DSN")
	flag.Parse()

	db, err := initMySQLDB(app.config.dsn)

	if err != nil {
		log.Panic(err)
	}

	srv := &http.Server{
		Addr:              port,
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
	}

	app.App = configuration.New(db)

	fmt.Println("Server listening on port", port)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("Someting went wrong to start the server")
	}

}
