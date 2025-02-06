package bootstrap

import (
	"context"
	"time"

	"github.com/korroziea/photo-storage/internal/config"
	"github.com/korroziea/photo-storage/internal/handler"
	httpserver "github.com/korroziea/photo-storage/internal/server/http"
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
) *App {
	handler := handler.InitRoutes()
	
	srv := httpserver.New(cfg, handler)
	
	app := &App{
		l:   l,
		cfg: cfg,
		srv: srv,
	}

	return app
}

func (a *App) Run(ctx context.Context) {
	go func(){
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
