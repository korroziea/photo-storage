package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/korroziea/photo-storage/internal/bootstrap"
	"github.com/korroziea/photo-storage/internal/config"
	"github.com/sethvargo/go-envconfig"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()
	quitSignal := make(chan os.Signal, 1)
	signal.Notify(quitSignal, os.Interrupt)

	var cfg config.Config
	if err := envconfig.Process(ctx, &cfg); err != nil {
		log.Fatal(err)
	}

	l, _ := zap.NewProduction()
	defer l.Sync()
	l.Info("logger initialiazed")

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	app, err := bootstrap.New(*l, cfg) // pointer
	if err != nil {
		l.Error("bootstrap.New: %w", zap.Error(err)) // TODO tut
	}

	go func(){
		osCall := <-quitSignal
		log.Printf("\nsystem call: %+v", osCall)
		cancel()
	}()

	app.Run(ctx)
}
