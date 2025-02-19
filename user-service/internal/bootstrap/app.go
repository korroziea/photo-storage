package bootstrap

import (
	"context"
	"fmt"
	"time"

	"github.com/korroziea/photo-storage/internal/config"
	"github.com/korroziea/photo-storage/internal/handler"
	userhndl "github.com/korroziea/photo-storage/internal/handler/user"
	"github.com/korroziea/photo-storage/internal/repository/psql"
	userrepo "github.com/korroziea/photo-storage/internal/repository/psql/user"
	httpserver "github.com/korroziea/photo-storage/internal/server/http"
	userservice "github.com/korroziea/photo-storage/internal/service/user"
	"github.com/korroziea/photo-storage/pkg/hashing"
	"go.uber.org/zap"
)

type App struct {
	l   zap.Logger
	cfg config.Config
	srv *httpserver.Server // pointer
}

func New(
	l zap.Logger,
	cfg config.Config,
) (*App, error) {
	db, _, err := psql.Connect(cfg.DB) // TODO add defer close
	if err != nil {
		return nil, fmt.Errorf("psql.Connect: %w", err)
	}
	
	argon := hashing.New(cfg.Hashing)
	
	userRepo := userrepo.New(db)

	userService := userservice.New(argon, userRepo)

	userHandler := userhndl.New(l, userService)

	handler := handler.New(userHandler)

	srv := httpserver.New(cfg, handler.InitRoutes())

	app := &App{
		l:   l,
		cfg: cfg,
		srv: srv,
	}

	return app, nil
}

func (a *App) Run(ctx context.Context) {
	go func() {
		if err := a.srv.ListenAndServer(); err != nil {
			a.l.Error("ListenAndServer", zap.Error(err))

			return
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	if err := a.srv.Shutdown(shutdownCtx); err != nil {
		a.l.Error("Shutdown", zap.Error(err))
	}
}
