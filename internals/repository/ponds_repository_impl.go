package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/masrayfa/go-delos-aqua/internals/models/domain"
)

type PondsRepositoryImpl struct {
}

func NewPondsRepositoryImpl() PondsRepository {
	return &PondsRepositoryImpl{}
}

func (p *PondsRepositoryImpl) FindAll(ctx context.Context, dbpool *pgxpool.Pool) ([]domain.Pond, error) {
	rows, err := dbpool.Query(ctx, "SELECT * FROM ponds")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ponds []domain.Pond
	for rows.Next() {
		var pond domain.Pond
		err := rows.Scan(&pond.PondId, &pond.Name)
		if err != nil {
			return nil, err
		}
		ponds = append(ponds, pond)
	}

	return ponds, nil
}

func (p *PondsRepositoryImpl) FindById(ctx context.Context, dbpool *pgxpool.Pool, id int) (domain.Pond, error) {
	var pond domain.Pond
	err := dbpool.QueryRow(ctx, "SELECT * FROM ponds WHERE pond_id = $1", id).Scan(&pond.PondId, &pond.Name)
	if err != nil {
		return domain.Pond{}, err
	}

	return pond, nil
}

func (p *PondsRepositoryImpl) Create(ctx context.Context, dbpool *pgxpool.Pool, pond domain.Pond) (domain.Pond, error) {
	err := dbpool.QueryRow(ctx, "INSERT INTO ponds (name, owner) VALUES ($1, $2) RETURNING pond_id", pond.Name).Scan(&pond.PondId)
	if err != nil {
		return domain.Pond{}, err
	}

	return pond, nil
}

func (p *PondsRepositoryImpl) Update(ctx context.Context, dbpool *pgxpool.Pool, pond domain.Pond) (domain.Pond, error) {
	_, err := dbpool.Exec(ctx, "UPDATE ponds SET name = $1, owner = $2 WHERE pond_id = $3", pond.Name, pond.PondId)
	if err != nil {
		return domain.Pond{}, err
	}

	return pond, nil
}

func (p *PondsRepositoryImpl) Delete(ctx context.Context, dbpool *pgxpool.Pool, id int) error {
	_, err := dbpool.Exec(ctx, "DELETE FROM ponds WHERE pond_id = $1", id)
	if err != nil {
		return err
	}

	return nil
}