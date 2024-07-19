package repository

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/masrayfa/go-delos-aqua/internals/helper"
	"github.com/masrayfa/go-delos-aqua/internals/models/domain"
)

type FarmRepositoryImpl struct {
}

func NewFarmRepository() FarmRepository {
	return &FarmRepositoryImpl{}
}

func (f *FarmRepositoryImpl) FindAll(ctx context.Context, dbpool *pgxpool.Pool) ([]domain.Farm, error) {
	log.Println("@FarmRepositoryImpl.FindAll:start")

	tx, err := dbpool.Begin(ctx)
	if err != nil {
		log.Println("@FarmRepositoryImpl.FindAll -> dbpool.Begin-error: ", err)
		return nil, err
	}
	defer helper.CommitOrRollback(ctx, tx)

	rows, err := tx.Query(ctx, "SELECT farm_id, user_id, name, location FROM farms WHERE deleted_at IS NULL")
	if err != nil {
		log.Println("@FarmRepositoryImpl.FindAll -> tx.Query-error: ", err)
		return nil, err
	}
	defer rows.Close()

	var farms []domain.Farm
	for rows.Next() {
		var farm domain.Farm
		err := rows.Scan(&farm.FarmId, &farm.UserId, &farm.Name, &farm.Location)
		if err != nil {
			log.Println("@FarmRepositoryImpl.FindAll -> rows.Scan-error: ", err)
			return nil, err
		}
		farms = append(farms, farm)
	}

	log.Println("@FarmRepositoryImpl.FindAll:succeed", farms)
	return farms, nil
}

func (f *FarmRepositoryImpl) FindById(ctx context.Context, dbpool *pgxpool.Pool, id int) (domain.Farm, error) {
	log.Println("@FarmRepositoryImpl.FindById:start")

	tx, err := dbpool.Begin(ctx)
	if err != nil {
		log.Println("@FarmRepositoryImpl.FindById -> dbpool.Begin-error: ", err)
		return domain.Farm{}, err
	}
	defer helper.CommitOrRollback(ctx, tx)

	var farm domain.Farm
	err = tx.QueryRow(ctx, "SELECT farm_id, user_id, name, location, created_at, updated_at, deleted_at FROM farms WHERE farm_id = $1 AND deleted_at IS NULL", id).Scan(&farm.FarmId, &farm.UserId, &farm.Name, &farm.Location)
	if err != nil {
		log.Println("@FarmRepositoryImpl.FindById -> tx.QueryRow-error: ", err)
		return domain.Farm{}, err
	}

	log.Println("@FarmRepositoryImpl.FindById:succeed", farm)
	return farm, nil
}

func (f *FarmRepositoryImpl) Create(ctx context.Context, dbpool *pgxpool.Pool, farm domain.Farm) (domain.Farm, error) {
	log.Println("@FarmRepositoryImpl.Create:start")

	tx, err := dbpool.Begin(ctx)
	if err != nil {
		log.Println("@FarmRepositoryImpl.Create -> dbpool.Begin-error: ", err)
		return domain.Farm{}, err
	}
	defer helper.CommitOrRollback(ctx, tx)

	_, err = tx.Exec(ctx, "INSERT INTO farms (user_id, name, location) VALUES ($1, $2, $3)", farm.UserId, farm.Name, farm.Location)
	if err != nil {
		log.Println("@FarmRepositoryImpl.Create -> tx.Exec-error: ", err)
		return domain.Farm{}, err
	}

	log.Println("@FarmRepositoryImpl.Create:succeed", farm)
	return farm, nil
}

func (f *FarmRepositoryImpl) Update(ctx context.Context, dbpool *pgxpool.Pool, farm domain.Farm) (domain.Farm, error) {
	log.Println("@FarmRepositoryImpl.Update:start")

	tx, err := dbpool.Begin(ctx)
	if err != nil {
		log.Println("@FarmRepositoryImpl.Update -> dbpool.Begin-error: ", err)
		return domain.Farm{}, err
	}
	defer helper.CommitOrRollback(ctx, tx)

	_, err = tx.Exec(ctx, "UPDATE farms SET name = $1, location = $2, updated_at = CURRENT_TIMESTAMP WHERE farm_id = $3 AND deleted_at IS NULL", farm.Name, farm.Location, farm.FarmId)
	if err != nil {
		log.Println("@FarmRepositoryImpl.Update -> tx.Exec-error: ", err)
		return domain.Farm{}, err
	}

	log.Println("@FarmRepositoryImpl.Update:succeed", farm)
	return farm, nil
}

func (f *FarmRepositoryImpl) Delete(ctx context.Context, dbpool *pgxpool.Pool, id int) error {
	log.Println("@FarmRepositoryImpl.Delete:start")

	tx, err := dbpool.Begin(ctx)
	if err != nil {
		log.Println("@FarmRepositoryImpl.Delete -> dbpool.Begin-error: ", err)
		return err
	}
	defer helper.CommitOrRollback(ctx, tx)

	_, err = tx.Exec(ctx, "UPDATE farms SET deleted_at = CURRENT_TIMESTAMP WHERE farm_id = $1", id)
	if err != nil {
		log.Println("@FarmRepositoryImpl.Delete -> tx.Exec-error: ", err)
		return err
	}

	log.Println("@FarmRepositoryImpl.Delete:succeed")
	return nil
}
