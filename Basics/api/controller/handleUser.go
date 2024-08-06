package controller

import (
	"api/models"
	"encoding/json"
	"net/http"
)

func HandleUser(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(getUser())

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func getUser() []models.User {
	return *models.FakeUsers()
}
