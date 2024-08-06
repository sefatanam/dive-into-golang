package main

import (
	"api/controller"
	"api/middleware"
	"log"
	"net/http"
)

func main() {
	http.Handle("/user", middleware.LoggingMiddleware(http.HandlerFunc(controller.HandleUser)))
	http.Handle("/todo", middleware.LoggingMiddleware(middleware.AuthorizationMiddleware(http.HandlerFunc(controller.HandleTodo))))

	log.Fatal(http.ListenAndServe("localhost:3000", nil))

}
