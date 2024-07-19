package service

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/masrayfa/go-delos-aqua/internals/models/domain"
	"github.com/masrayfa/go-delos-aqua/internals/models/web"
	"github.com/masrayfa/go-delos-aqua/internals/repository"
)

type PondsServiceImpl struct {
	pondsRepository repository.PondsRepository
	db *pgxpool.Pool
}

func NewPondsService(pondsRepository repository.PondsRepository, db *pgxpool.Pool) PondsService {
	return &PondsServiceImpl{
		pondsRepository: pondsRepository,
		db: db,
	}
}

func (ps *PondsServiceImpl) FindAll(ctx context.Context) ([]web.PondResponse, error) {
	ponds, err := ps.pondsRepository.FindAll(ctx, ps.db)
	if err != nil {
		return nil, err
	}

	var pondRes []web.PondResponse
	for _, pond := range ponds {
		pondRes = append(pondRes, web.PondResponse{
			PondId: pond.PondId,
			Name: pond.Name,
		})
	}

	return pondRes, nil
}

func (ps *PondsServiceImpl) FindById(ctx context.Context, id int) (web.PondResponse, error) {
	pond, err := ps.pondsRepository.FindById(ctx, ps.db, id)
	if err != nil {
		return web.PondResponse{}, err
	}

	pondRes := web.PondResponse{
		PondId: pond.PondId,
		Name: pond.Name,
	}

	return pondRes, nil
}

func (ps *PondsServiceImpl) Create(ctx context.Context, payload web.PondCreateRequest) (web.PondResponse, error) {
	pondDomain := domain.Pond{
		Name: payload.Name,
	}

	pond, err := ps.pondsRepository.Create(ctx, ps.db, pondDomain)
	if err != nil {
		return web.PondResponse{}, err
	}

	pondRes := web.PondResponse{
		PondId: pond.PondId,
		Name: pond.Name,
	}

	return pondRes, nil
}

func (ps *PondsServiceImpl) Update(ctx context.Context, payload web.PondUpdateRequest, id int) (web.PondResponse, error) {
	pondDomain := domain.Pond{
		PondId: id,
		Name: payload.Name,
	}

	pond, err := ps.pondsRepository.Update(ctx, ps.db, pondDomain)
	if err != nil {
		return web.PondResponse{}, err
	}

	pondRes := web.PondResponse{
		PondId: pond.PondId,
		Name: pond.Name,
	}

	return pondRes, nil
}

func (ps *PondsServiceImpl) Delete(ctx context.Context, id int) error {
	err := ps.pondsRepository.Delete(ctx, ps.db, id)
	if err != nil {
		return err
	}

	return nil
}