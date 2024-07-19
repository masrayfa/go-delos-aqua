package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/masrayfa/go-delos-aqua/internals/models/domain"
)

type UserRepository interface {
	FindAll(ctx context.Context, dbpool *pgxpool.Pool) ([]domain.User, error)
	FindById(ctx context.Context, dbpool *pgxpool.Pool, id int) (domain.User, error)
	Create(ctx context.Context, dbpool *pgxpool.Pool, user domain.User) (domain.User, error)
	Update(ctx context.Context, dbpool *pgxpool.Pool, user domain.User, id int) (domain.User, error)
	Delete(ctx context.Context, dbpool *pgxpool.Pool, id int) error
}