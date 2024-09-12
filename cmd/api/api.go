package api

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/roman-haidarov/go-rest-api/cmd/service/client"
	"net/http"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	clientHandler := client.NewHandler()
	clientHandler.RegisterRoutes(subrouter)

	return http.ListenAndServe(s.addr, router)
}
