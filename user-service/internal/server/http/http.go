package http

import (
	"context"
	"net/http"
	"time"

	"github.com/korroziea/photo-storage/internal/config"
)

const (
	readHeaderTimeout = 10 * time.Second
	readTimeout       = 10 * time.Second
	writeTimeout      = time.Minute
)

type Server struct {
	server http.Server
}

func New(cfg config.Config, handler http.Handler) *Server {
	s := &Server{
		server: http.Server{
			Addr:              ":" + cfg.HTTPPort,
			Handler:           handler,
			ReadHeaderTimeout: readHeaderTimeout,
			ReadTimeout:       readTimeout,
			WriteTimeout:      writeTimeout,
		},
	}

	return s
}

func (s *Server) ListenAndServer() error {
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
