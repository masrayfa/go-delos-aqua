package service

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/masrayfa/go-delos-aqua/internals/models/domain"
	"github.com/masrayfa/go-delos-aqua/internals/models/web"
	"github.com/masrayfa/go-delos-aqua/internals/repository"
)

type FarmServiceImpl struct {
	farmRepository repository.FarmRepository
	db 		   *pgxpool.Pool
}

func NewFarmService(farmRepository repository.FarmRepository, db *pgxpool.Pool) FarmService {
	return &FarmServiceImpl{
		farmRepository: farmRepository,
		db: db,
	}
}

func (fs *FarmServiceImpl) FindAll(ctx context.Context) ([]web.FarmRead, error) {
	farms, err := fs.farmRepository.FindAll(ctx, fs.db)
	if err != nil {
		return nil, err
	}

	var farmReads []web.FarmRead
	for _, farm := range farms {
		farmReads = append(farmReads, web.FarmRead{
			FarmId: farm.FarmId,
			Name: farm.Name,
			Location: farm.Location,
			Owner: farm.Owner,
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
		Name: farm.Name,
		Location: farm.Location,
		Owner: farm.Owner,
	}

	return farmRead, nil
}

func (fs *FarmServiceImpl) Create(ctx context.Context, payload web.FarmRequest) (web.FarmRead, error) {
	farmDomain := domain.Farm{
		Name: payload.Name,
		Location: payload.Location,
		Owner: payload.Owner,
	}

	farm, err := fs.farmRepository.Create(ctx, fs.db, farmDomain)
	if err != nil {
		return web.FarmRead{}, err
	}

	farmRead := web.FarmRead{
		FarmId: farm.FarmId,
		Name: farm.Name,
		Location: farm.Location,
		Owner: farm.Owner,
	}

	return farmRead, nil
}

func (fs *FarmServiceImpl) Update(ctx context.Context, payload web.FarmRequest) (web.FarmRead, error) {
	farmDomain := domain.Farm{
		Name: payload.Name,
		Location: payload.Location,
		Owner: payload.Owner,
	}

	farm, err := fs.farmRepository.Update(ctx, fs.db, farmDomain)
	if err != nil {
		return web.FarmRead{}, err
	}

	farmRead := web.FarmRead{
		FarmId: farm.FarmId,
		Name: farm.Name,
		Location: farm.Location,
		Owner: farm.Owner,
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
