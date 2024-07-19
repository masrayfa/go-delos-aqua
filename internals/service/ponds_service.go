package service

import (
	"context"

	"github.com/masrayfa/go-delos-aqua/internals/models/web"
)

type PondsService interface {
	FindAll(ctx context.Context) ([]web.PondResponse, error)
	FindById(ctx context.Context, id int) (web.PondResponse, error)
	Create(ctx context.Context, payload web.PondCreateRequest) (web.PondResponse, error)
	Update(ctx context.Context, payload web.PondUpdateRequest, id int) (web.PondResponse, error)
	Delete(ctx context.Context, id int) error
}