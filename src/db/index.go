package db

import (
	"context"
	"fmt"
	"sync"
	"time"
	"todo_list/src/conf"
	"todo_list/src/db/repos"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	Pool *pgxpool.Pool
	Q    *repos.Queries
	once sync.Once
)

func Connect() {
	once.Do(func() {
		cfg, err := pgxpool.ParseConfig(conf.Env.DB_URL)
		if err != nil {
			panic(fmt.Sprintf("❌ Invalid DB_URL: %v", err))
		}
		// Pool config
		cfg.MaxConns = 25
		cfg.MinConns = 5
		cfg.MaxConnLifetime = 30 * time.Minute
		cfg.MaxConnIdleTime = 10 * time.Minute
		cfg.HealthCheckPeriod = 1 * time.Minute
		// Create pool
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		Pool, err = pgxpool.NewWithConfig(ctx, cfg)
		if err != nil {
			panic(fmt.Sprintf("❌ Unable to connect to database: %v", err))
		}
		if err := Pool.Ping(ctx); err != nil {
			panic(fmt.Sprintf("❌ Database ping failed: %v", err))
		}
		// Add query
		Q = repos.New(Pool)
	})
}

func Disconnect() {
	if Pool != nil {
		Pool.Close()
		fmt.Println("✅ DB connection closed!")
	}
}
