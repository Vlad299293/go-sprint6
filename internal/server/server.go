package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

const (
	ServerAddr   = ":8080"
	ReadTimeout  = 5 * time.Second
	WriteTimeout = 10 * time.Second
	IdleTimeout  = 15 * time.Second
)

type Server struct {
	log *log.Logger
	srv *http.Server
}

func NewServer(logger *log.Logger) *Server {
	router := http.NewServeMux()

	router.HandleFunc("/", handlers.IndexHandler)
	router.HandleFunc("/upload", handlers.UploadHandler)

	httpServer := &http.Server{
		Addr:         ServerAddr,
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  ReadTimeout,
		WriteTimeout: WriteTimeout,
		IdleTimeout:  IdleTimeout,
	}

	return &Server{
		log: logger,
		srv: httpServer,
	}
}

func (s *Server) Run() error {
	return s.srv.ListenAndServe()
}
