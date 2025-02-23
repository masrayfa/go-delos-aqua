package repository

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/masrayfa/go-delos-aqua/internals/helper"
	"github.com/masrayfa/go-delos-aqua/internals/models/domain"
)

type PondsRepositoryImpl struct {
}

func NewPondsRepository() PondsRepository {
	return &PondsRepositoryImpl{}
}

func (p *PondsRepositoryImpl) FindAll(ctx context.Context, dbpool *pgxpool.Pool) ([]domain.Pond, error) {
	log.Println("@PondsRepositoryImpl.FindAll:start")

	tx, err := dbpool.Begin(ctx)
	if err != nil {
		log.Println("@PondsRepositoryImpl.FindAll -> dbpool.Begin-error: ", err)
		return nil, err
	}

	rows, err := tx.Query(ctx, "SELECT pond_id, farm_id, name FROM ponds")
	if err != nil {
		log.Println("@PondsRepositoryImpl.FindAll -> dbpool.Query-error: ", err)
		return nil, err
	}

	var ponds []domain.Pond
	for rows.Next() {
		var pond domain.Pond
		err := rows.Scan(&pond.PondId, &pond.FarmId, &pond.Name)
		if err != nil {
			log.Println("@PondsRepositoryImpl.FindAll -> rows.Scan-error: ", err)
			return nil, err
		}
		ponds = append(ponds, pond)
	}

	defer rows.Close()

	log.Println("@PondsRepositoryImpl.FindAll:succeed")
	return ponds, nil
}

func (p *PondsRepositoryImpl) FindById(ctx context.Context, dbpool *pgxpool.Pool, id int) (domain.Pond, error) {
	log.Println("@PondsRepositoryImpl.FindById:start")

	tx, err := dbpool.Begin(ctx)
	if err != nil {
		log.Println("@PondsRepositoryImpl.FindById -> dbpool.Begin-error: ", err)
		return domain.Pond{}, err
	}

	var pond domain.Pond
	err = tx.QueryRow(ctx, "SELECT pond_id, farm_id, name FROM ponds WHERE pond_id = $1", id).Scan(&pond.PondId, &pond.FarmId, &pond.Name)
	if err != nil {
		log.Println("@PondsRepositoryImpl.FindById -> dbpool.QueryRow-error: ", err)
		return domain.Pond{}, err
	}

	log.Println("@PondsRepositoryImpl.FindById:succeed")
	return pond, nil
}

func (p *PondsRepositoryImpl) Create(ctx context.Context, dbpool *pgxpool.Pool, pond domain.Pond) (domain.Pond, error) {
	log.Println("@PondsRepositoryImpl.Create:start")

	tx, err := dbpool.Begin(ctx)
	if err != nil {
		log.Println("@PondsRepositoryImpl.Create -> dbpool.Begin-error: ", err)
		return domain.Pond{}, err
	}

	defer helper.CommitOrRollback(ctx, tx)
	// defer func() {
	// 	if err != nil {
	// 		log.Println("@PondsRepositoryImpl.Create -> rollback due to error: ", err)
	// 		if rbErr := tx.Rollback(ctx); rbErr != nil {
	// 			log.Printf("tx.Rollback failed: %v\n", rbErr)
	// 		}
	// 	} else {
	// 		if cmErr := tx.Commit(ctx); cmErr != nil {
	// 			log.Printf("tx.Commit failed: %v\n", cmErr)
	// 		}
	// 	}
	// }()

	log.Println("@PondsRepositoryImpl.Create:pond", pond)
	err = tx.QueryRow(ctx, "INSERT INTO ponds (farm_id, name) VALUES ($1, $2) RETURNING pond_id", pond.FarmId, pond.Name).Scan(&pond.PondId)
	if err != nil {
		log.Println("@PondsRepositoryImpl.Create -> dbpool.QueryRow-error: ", err)
		return domain.Pond{}, err
	}

	log.Println("@PondsRepositoryImpl.Create:succeed")
	return pond, nil
}

func (p *PondsRepositoryImpl) Update(ctx context.Context, dbpool *pgxpool.Pool, pond domain.Pond) (domain.Pond, error) {
	log.Println("@PondsRepositoryImpl.Update:start")

	tx, err := dbpool.Begin(ctx)
	if err != nil {
		log.Println("@PondsRepositoryImpl.Update -> dbpool.Begin-error: ", err)
		return domain.Pond{}, err
	}
	defer helper.CommitOrRollback(ctx, tx)
	// defer func() {
	// 	if err != nil {
	// 		log.Println("@PondsRepositoryImpl.Create -> rollback due to error: ", err)
	// 		if rbErr := tx.Rollback(ctx); rbErr != nil {
	// 			log.Printf("tx.Rollback failed: %v\n", rbErr)
	// 		}
	// 	} else {
	// 		if cmErr := tx.Commit(ctx); cmErr != nil {
	// 			log.Printf("tx.Commit failed: %v\n", cmErr)
	// 		}
	// 	}
	// }()

	log.Println("@PondsRepositoryImpl.Update:pond", pond)
	_, err = tx.Exec(ctx, "UPDATE ponds SET name = $1, farm_id = $2 WHERE pond_id = $3", pond.Name, pond.FarmId, pond.PondId)
	if err != nil {
		log.Println("@PondsRepositoryImpl.Update -> dbpool.Exec-error: ", err)
		return domain.Pond{}, err
	}


	log.Println("@PondsRepositoryImpl.Update:succeed")
	return pond, nil
}

func (p *PondsRepositoryImpl) Delete(ctx context.Context, dbpool *pgxpool.Pool, id int) error {
	log.Println("@PondsRepositoryImpl.Delete:start")

	tx, err := dbpool.Begin(ctx)
	if err != nil {
		log.Println("@PondsRepositoryImpl.Delete -> dbpool.Begin-error: ", err)
		return err
	}

	_, err = tx.Exec(ctx, "DELETE FROM ponds WHERE pond_id = $1", id)
	if err != nil {
		log.Println("@PondsRepositoryImpl.Delete -> dbpool.Exec-error: ", err)
		return err
	}

	log.Println("@PondsRepositoryImpl.Delete:succeed")
	return nil
}