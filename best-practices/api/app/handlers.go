package app

import (
	"api/utils"
	"encoding/json"
	"net/http"
)

func (s *Server) handleGetUserById(w http.ResponseWriter, r *http.Request) {
	user := s.store.Get(1)
	json.NewEncoder(w).Encode(user)
}

func (s *Server) handleDeleteUserById(w http.ResponseWriter, r *http.Request) {
	user := s.store.Get(1)
	_ = utils.Round2Dec(10.234)
	json.NewEncoder(w).Encode(user)
}
