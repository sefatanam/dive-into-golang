package util

import (
	"encoding/json"
	"net/http"
)

func SendData(w http.ResponseWriter, data any, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(data); err != nil {
		http.Error(w, err.Error(), 400)
	}
}
