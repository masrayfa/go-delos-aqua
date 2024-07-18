package service

import (
	"context"

	"github.com/masrayfa/go-delos-aqua/internals/models/web"
)

type UserService interface {
	FindAll(ctx context.Context) ([]web.UserRead, error)
	FindById(ctx context.Context, id int) (web.UserRead, error)
	Create(ctx context.Context, payload web.UserCreate) (web.UserRead, error)
	Update(ctx context.Context, payload web.UserUpdate) (web.UserRead, error)
	Delete(ctx context.Context, id int) error
}