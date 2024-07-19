package repository

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/masrayfa/go-delos-aqua/internals/helper"
	"github.com/masrayfa/go-delos-aqua/internals/models/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (ur *UserRepositoryImpl) FindAll(ctx context.Context, dbpool *pgxpool.Pool) ([]domain.User, error) {
	log.Println("@UserRepositoryImpl.FindAll:start")

	tx, err := dbpool.Begin(ctx)
	if err != nil {
		log.Println("@UserRepositoryImpl.FindAll -> dbpool.Begin: ", err)
		return nil, err
	}
	defer helper.CommitOrRollback(ctx, tx)

	rows, err := tx.Query(ctx, "SELECT user_id, email, username FROM users WHERE deleted_at IS NULL")
	if err != nil {
		log.Println("@UserRepositoryImpl.FindAll -> dbpool.Query: ", err)
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.UserId, &user.Email, &user.Username)
		if err != nil {
			log.Println("@UserRepositoryImpl.FindAll -> rows.Scan: ", err)
			return nil, err
		}
		users = append(users, user)
	}

	log.Println("@UserRepositoryImpl.FindAll:succeed")
	return users, nil
}

func (ur *UserRepositoryImpl) FindById(ctx context.Context, dbpool *pgxpool.Pool, id int) (domain.User, error) {
	log.Println("@UserRepositoryImpl.FindById:start")

	var user domain.User

	tx, err := dbpool.Begin(ctx)
	if err != nil {
		log.Println("@UserRepositoryImpl.FindById:error: ", err)
		return domain.User{}, err
	}
	defer helper.CommitOrRollback(ctx, tx)

	err = tx.QueryRow(ctx, "SELECT user_id, email, username FROM users WHERE user_id = $1 AND deleted_at IS NULL", id).Scan(&user.UserId, &user.Email, &user.Username)
	if err != nil {
		log.Println("@UserRepositoryImpl.FindById:error: ", err)
		return domain.User{}, err
	}

	log.Println("@UserRepositoryImpl.FindById:succeed")
	return user, nil
}

func (ur *UserRepositoryImpl) Create(ctx context.Context, dbpool *pgxpool.Pool, user domain.User) (domain.User, error) {
	log.Println("@UserRepositoryImpl.Create:start")

	tx, err := dbpool.Begin(ctx)
	if err != nil {
		log.Println("@UserRepositoryImpl.Create:error: ", err)
		return domain.User{}, err
	}
	defer helper.CommitOrRollback(ctx, tx)

	err = tx.QueryRow(ctx, "INSERT INTO users (email, username, password) VALUES ($1, $2, $3) RETURNING user_id", user.Email, user.Username, user.Password).Scan(&user.UserId)
	if err != nil {
		log.Println("@UserRepositoryImpl.Create:error: ", err)
		return domain.User{}, err
	}

	log.Println("@UserRepositoryImpl.Create:succeed")
	return user, nil
}

func (ur *UserRepositoryImpl) Update(ctx context.Context, dbpool *pgxpool.Pool, user domain.User, id int) (domain.User, error) {
	log.Println("@UserRepositoryImpl.Update:start")

	tx, err := dbpool.Begin(ctx)
	if err != nil {
		log.Println("@UserRepositoryImpl.Update:error: ", err)
		return domain.User{}, err
	}
	defer helper.CommitOrRollback(ctx, tx)

	_, err = tx.Exec(ctx, "UPDATE users SET email = $1, username = $2 WHERE user_id = $3 AND deleted_at IS NULL", user.Email, user.Username, id) 
	if err != nil {
		log.Println("@UserRepositoryImpl.Update:error: ", err)
		return domain.User{}, err
	}

	log.Println("@UserRepositoryImpl.Update:succeed")
	return user, nil
}

func (ur *UserRepositoryImpl) Delete(ctx context.Context, dbpool *pgxpool.Pool, id int) error {
	log.Println("@UserRepositoryImpl.Delete:start")

	tx, err := dbpool.Begin(ctx)
	if err != nil {
		log.Println("@UserRepositoryImpl.Delete:error: ", err)
		return err
	}
	defer helper.CommitOrRollback(ctx, tx)

	_, err = tx.Exec(ctx, "UPDATE users SET deleted_at = CURRENT_TIMESTAMP WHERE user_id = $1", id)
	if err != nil {
		log.Println("@UserRepositoryImpl.Delete:error: ", err)
		return err
	}

	log.Println("@UserRepositoryImpl.Delete:succeed")
	return nil
}
