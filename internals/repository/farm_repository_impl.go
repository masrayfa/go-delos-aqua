package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/masrayfa/go-delos-aqua/internals/models/domain"
)

type FarmRepositoryImpl struct {
}

func NewFarmRepositoryImpl() FarmRepository {
	return &FarmRepositoryImpl{}
}

func (f *FarmRepositoryImpl) FindAll(ctx context.Context, dbpool *pgxpool.Pool) ([]domain.Farm, error) {
	panic("implement me")
}

func (f *FarmRepositoryImpl) FindById(ctx context.Context, dbpool *pgxpool.Pool, id int) (domain.Farm, error) {
	panic("implement me")
}

func (f *FarmRepositoryImpl) Create(ctx context.Context, dbpool *pgxpool.Pool, farm domain.Farm) (domain.Farm, error) {
	panic("implement me")
}

func (f *FarmRepositoryImpl) Update(ctx context.Context, dbpool *pgxpool.Pool, farm domain.Farm) (domain.Farm, error) {
	panic("implement me")
}

func (f *FarmRepositoryImpl) Delete(ctx context.Context, dbpool *pgxpool.Pool, id int) error {
	panic("implement me")
}
