package service

import (
	"context"

	"github.com/masrayfa/go-delos-aqua/internals/models/web"
)

type FarmService interface {
	FindAll(ctx context.Context) ([]web.FarmRead, error)
	FindById(ctx context.Context, id int) (web.FarmRead, error)
	Create(ctx context.Context, payload web.FarmRequest) (web.FarmRead, error)
	Update(ctx context.Context, payload web.FarmRequest) (web.FarmRead, error)
	Delete(ctx context.Context, id int) error
}