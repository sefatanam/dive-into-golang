package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

const keyServerAddr = "serverAddr"

func getRoot(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("got /request\n")
	_, err := io.WriteString(w, "This is my website.\n")
	if err != nil {
		return
	}
}

func getHello(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("got hello /request\n")
	_, err := io.WriteString(w, "Hello HTTP!\n")

	if err != nil {
		return
	}

}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)

	err := http.ListenAndServe(":3000", mux)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server:%s\n", err)
		os.Exit(1)
	}
}
