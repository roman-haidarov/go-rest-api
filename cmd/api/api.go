package api

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"github.com/roman-haidarov/go-rest-api/cmd/service/client"
)

type APIServer struct {
	addr       string
	db         *sql.DB
	httpServer *http.Server
	router     *mux.Router
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	server := &APIServer{
		addr:   addr,
		db:     db,
		router: mux.NewRouter(),
	}
	server.httpServer = &http.Server{
		Addr:           addr,
		Handler:        server.router,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return server
}

func (s *APIServer) SetupRoutes() {
	subrouter := s.router.PathPrefix("/api/v1").Subrouter()
	clientStore := client.NewStore(s.db)
	clientHandler := client.NewHandler(clientStore)
	clientHandler.RegisterRoutes(subrouter)
}

func (s *APIServer) Run() error {
	s.SetupRoutes()
	log.Printf("Starting server on %s", s.addr)
	return s.httpServer.ListenAndServe()
}

func (s *APIServer) Shutdown(ctx context.Context) error {
	log.Println("Server is shutting down...")
	
	err := s.httpServer.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("server shutdown failed: %v", err)
	}
	
	log.Println("Server stopped")
	return nil
}

func (s *APIServer) GracefulShutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	time.Sleep(500 * time.Millisecond)

	if err := s.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}

	if err := s.db.Close(); err != nil {
		log.Printf("Error closing database connection: %v", err)
	}
}
