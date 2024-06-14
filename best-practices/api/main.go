package main

import (
	"api/app"
	"api/storage"
	"flag"
	"log"
)

func main() {
	listenAddr := flag.String("listenAddr", ":3000", "the server address")
	flag.Parse()

	storage := storage.NewMemoryStorage()
	flag.Parse()

	app := app.NewServer(*listenAddr, storage)
	log.Println("server running on port", *listenAddr)
	log.Fatal(app.Start())
}
