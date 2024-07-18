package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/masrayfa/go-delos-aqua/internals/models/domain"
)

type FarmRepository interface {
	FindAll(ctx context.Context, dbpool *pgxpool.Pool) ([]domain.Farm, error)
	FindById(ctx context.Context, dbpool *pgxpool.Pool, id int) (domain.Farm, error)
	Create(ctx context.Context, dbpool *pgxpool.Pool, farm domain.Farm) (domain.Farm, error)
	Update(ctx context.Context, dbpool *pgxpool.Pool, farm domain.Farm) (domain.Farm, error)
	Delete(ctx context.Context, dbpool *pgxpool.Pool, id int) error
}
