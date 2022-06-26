package rest

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"ms-hw/internal/config"
	"net/http"
	"time"
)

type Server struct {
	cfg *config.Cfg
	srv *http.Server
}

func NewServer(cfg *config.Cfg) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) Serve(routes *mux.Router) {
	log.Println("Starting server")

	s.srv = &http.Server{
		Addr:    ":" + s.cfg.HttpPort,
		Handler: routes,
	}

	log.Println("HW server started")
	err := s.srv.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		log.Println(err)
	}
}

func (s *Server) Shutdown() {
	log.Printf("server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := s.srv.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("server Shutdown Failed:%s", err)
	}

	log.Printf("server has been terminated")
}
