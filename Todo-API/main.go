package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type todo struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{Id: "1", Title: "Complete Struct in GO", Completed: false},
	{Id: "2", Title: "Complete Pointer in GO", Completed: false},
	{Id: "3", Title: "Complete REST API in GO", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newTodo todo
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}
	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}

	context.IndentedJSON(http.StatusFound, todo)

}

func toggleStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}

	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)

}

func getById(id string) (*todo, error) {
	for i, t := range todos {
		if t.Id == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleStatus)
	router.POST("/todos", addTodo)
	err := router.Run("localhost:3000")
	if err != nil {
		return
	}
}
