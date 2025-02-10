package psql

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/korroziea/photo-storage/internal/config"
)

const connectionTimeout = 3 * time.Second

func Connect(cfg config.Postgres) (*pgxpool.Pool, func(), error) {
	config, err := pgxpool.ParseConfig(cfg.PostgresURL())
	if err != nil {
		return nil, nil, fmt.Errorf("pgxpool.ParseConfig: %w", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, config) // why do we need pool
	if err != nil {
		return nil, nil, fmt.Errorf("pgxpool.NewWithConfig: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, nil, fmt.Errorf("pool.Ping: %w", err)
	}

	return pool, pool.Close, nil
}
