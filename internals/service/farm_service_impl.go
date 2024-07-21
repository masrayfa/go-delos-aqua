package service

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/masrayfa/go-delos-aqua/internals/dependencies"
	"github.com/masrayfa/go-delos-aqua/internals/helper"
	"github.com/masrayfa/go-delos-aqua/internals/models/domain"
	"github.com/masrayfa/go-delos-aqua/internals/models/web"
	"github.com/masrayfa/go-delos-aqua/internals/repository"
)

type FarmServiceImpl struct {
	farmRepository repository.FarmRepository
	db 		   *pgxpool.Pool
	validate  *dependencies.Validator
}

func NewFarmService(farmRepository repository.FarmRepository, db *pgxpool.Pool, validate *dependencies.Validator) FarmService {
	return &FarmServiceImpl{
		farmRepository: farmRepository,
		db: db,
		validate: validate,
	}
}

func (fs *FarmServiceImpl) FindAll(ctx context.Context) ([]web.FarmRead, error) {
	farms, err := fs.farmRepository.FindAll(ctx, fs.db)
	if err != nil {
		return nil, err
	}

	if len(farms) == 0 {
		return nil, helper.ErrNotFound
	}

	var farmReads []web.FarmRead
	for _, farm := range farms {
		farmReads = append(farmReads, web.FarmRead{
			FarmId: farm.FarmId,
			UserId: farm.UserId,
			Name: farm.Name,
			Location: farm.Location,
		})
	}

	return farmReads, nil
}

func (fs *FarmServiceImpl) FindById(ctx context.Context, id int) (web.FarmRead, error) {
	farm, err := fs.farmRepository.FindById(ctx, fs.db, id)
	if err != nil {
		return web.FarmRead{}, err
	}

	farmRead := web.FarmRead{
		FarmId: farm.FarmId,
		UserId: farm.UserId,
		Name: farm.Name,
		Location: farm.Location,
	}

	return farmRead, nil
}

func (fs *FarmServiceImpl) Create(ctx context.Context, payload web.FarmRequest) (web.FarmRead, error) {
	err := fs.validate.ValidateStruct(payload)
	if err != nil {
		return web.FarmRead{}, err
	}

	farmDomain := domain.Farm{
		UserId: payload.UserId,
		Name: payload.Name,
		Location: payload.Location,
	}

	farm, err := fs.farmRepository.Create(ctx, fs.db, farmDomain)
	if err != nil {
		return web.FarmRead{}, err
	}

	farmRead := web.FarmRead{
		FarmId: farm.FarmId,
		Name: farm.Name,
		Location: farm.Location,
	}

	return farmRead, nil
}

func (fs *FarmServiceImpl) Update(ctx context.Context, payload web.FarmRequest) (web.FarmRead, error) {
	err := fs.validate.ValidateStruct(payload)
	if err != nil {
		return web.FarmRead{}, err
	}

	farmDomain := domain.Farm{
		Name: payload.Name,
		Location: payload.Location,
	}

	farm, err := fs.farmRepository.Update(ctx, fs.db, farmDomain)
	if err != nil {
		return web.FarmRead{}, err
	}

	farmRead := web.FarmRead{
		FarmId: farm.FarmId,
		UserId: farm.UserId,
		Name: farm.Name,
		Location: farm.Location,
	}

	return farmRead, nil
}

func (fs *FarmServiceImpl) Delete(ctx context.Context, id int) error {
	err := fs.farmRepository.Delete(ctx, fs.db, id)
	if err != nil {
		return err
	}

	return nil
}
