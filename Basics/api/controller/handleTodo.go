package controller

import (
	"api/models"
	"encoding/json"
	"net/http"
	"strconv"
)

func HandleTodo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		userId, isSuccess := parseUserId(r)
		if !isSuccess {
			http.Error(w, "Invalid user id.", http.StatusBadRequest)
			return
		}
		todos := getTodos(userId)
		json.NewEncoder(w).Encode(todos)

	case "POST":
		userId, isSuccess := parseUserId(r)
		if !isSuccess {
			http.Error(w, "Invalid user id.", http.StatusBadRequest)
			return
		}
		todo := models.Todo{Id: 11, Name: "Post", UserId: userId, Description: "Just now posted"}
		addTodo(todo)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(todo)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getTodos(userId int) *[]models.Todo {
	todos := []models.Todo{}
	for _, todo := range *models.FakeTodos() {
		shouldAppendTodo := todo.UserId == userId
		if shouldAppendTodo {
			todos = append(todos, todo)
		}
	}
	return &todos
}

func addTodo(todo models.Todo) bool {
	*models.FakeTodos() = append(*models.FakeTodos(), todo)
	return true
}

func parseUserId(r *http.Request) (int, bool) {
	userIdStr := r.Header.Get("userid")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return 0, false
	}
	return userId, true
}
