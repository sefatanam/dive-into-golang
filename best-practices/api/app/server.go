package app

import (
	"api/storage"
	"net/http"
)

type Server struct {
	listenAddr string
	store      storage.Storage
}

func NewServer(listenAddr string, store storage.Storage) *Server {
	return &Server{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/user", s.handleGetUserById)
	http.HandleFunc("/user/id", s.handleDeleteUserById)
	return http.ListenAndServe(s.listenAddr, nil)
}
