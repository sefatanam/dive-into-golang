package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

const keyServerAddr = "serverAddr"

func getRoot(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	hasFirst := r.URL.Query().Has("first")
	first := r.URL.Query().Get("first")

	fmt.Printf("%s got /request. first(%t)=%s \n", ctx.Value(keyServerAddr), hasFirst, first)
	_, err := io.WriteString(w, "This is my website.\n")
	if err != nil {
		return
	}
}

func getHello(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	fmt.Printf("%s got hello /request\n", ctx.Value(keyServerAddr))
	_, err := io.WriteString(w, "Hello HTTP!\n")

	if err != nil {
		return
	}

}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", getRoot)

	ctx, cancelCtx := context.WithCancel(context.Background())
	serverOne := &http.Server{
		Addr:    ":3001",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
	}

	serverTwo := &http.Server{
		Addr:    ":3003",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
	}

	go func() {
		err := serverOne.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server one closed\n")
		} else if err != nil {
			fmt.Printf("error listening for server one: %s\n", err)
		}
		cancelCtx()
	}()

	go func() {
		err := serverTwo.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server two closed\n")
		} else if err != nil {
			fmt.Printf("error listening for server two: %s\n", err)
		}
		cancelCtx()
	}()

	<-ctx.Done()

	mux.HandleFunc("/hello", getHello)

	err := http.ListenAndServe(":3000", mux)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server:%s\n", err)
		os.Exit(1)
	}
}
