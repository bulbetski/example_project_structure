package deps

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/samber/do"
)

func (c *Container) GetDatabase() *pgxpool.Pool {
	return do.MustInvoke[*pgxpool.Pool](c.i)
}

func (c *Container) provideDatabase() {
	do.Provide(c.i, func(*do.Injector) (*pgxpool.Pool, error) {
		cfg := c.GetConfig()
		pool, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
		if err != nil {
			return nil, fmt.Errorf("deps: create pool: %w", err)
		}
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err = pool.Ping(ctx); err != nil {
			pool.Close()
			return nil, fmt.Errorf("deps: ping database: %w", err)
		}
		return pool, nil
	})
}
