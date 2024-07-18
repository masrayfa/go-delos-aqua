package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/masrayfa/go-delos-aqua/internals/models/domain"
)

type PondsRepository interface {
	FindAll(ctx context.Context, dbpool *pgxpool.Pool) ([]domain.Pond, error)
	FindById(ctx context.Context, dbpool *pgxpool.Pool, id int) (domain.Pond, error)
	Create(ctx context.Context, dbpool *pgxpool.Pool, pond domain.Pond) (domain.Pond, error)
	Update(ctx context.Context, dbpool *pgxpool.Pool, pond domain.Pond) (domain.Pond, error)
	Delete(ctx context.Context, dbpool *pgxpool.Pool, id int) error
}
