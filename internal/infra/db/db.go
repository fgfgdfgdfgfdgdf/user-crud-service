package db

import (
	"context"
	"fmt"

	"github.com/fgfgdfgdfgfdgdf/user-crud-service/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPGPool(ctx context.Context) (*pgxpool.Pool, error) {
	connString := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s",
		config.POSTGRES.POSTGRES_USER, config.POSTGRES.POSTGRES_PASSWORD, config.POSTGRES.POSTGRES_HOST, config.POSTGRES.POSTGRES_PORT, config.POSTGRES.POSTGRES_DB,
	)

	return pgxpool.New(ctx, connString)
}
