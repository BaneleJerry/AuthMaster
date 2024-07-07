package server

import (
	"log"
	"net/http"

	"github.com/BaneleJerry/AuthMaster/internal/config"
	"github.com/gorilla/mux"
)

type Server struct {
	cfg *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

// Start starts the HTTP server and listens for incoming connections.
// It initializes a new Gorilla Mux router and registers the health endpoint.
// The server listens on the port specified in the configuration.
//
// Parameters:
// - None
//
// Returns:
// - error: An error if the server fails to start, or nil if successful.
func (s *Server) Start() error {
	router := mux.NewRouter()

	// Health endpoint
	// This endpoint is used to check the health of the server.
	// It responds with a 200 OK status code and the text "OK".
	router.HandleFunc("/health", s.handleHealth).Methods("GET")

	// Start server
	addr := ":" + s.cfg.Port
	log.Printf("Server starting on port %s", s.cfg.Port)
	return http.ListenAndServe(addr, router)
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
