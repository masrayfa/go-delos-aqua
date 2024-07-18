package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/masrayfa/go-delos-aqua/internals/models/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (ur *UserRepositoryImpl) FindAll(ctx context.Context, dbpool *pgxpool.Pool) ([]domain.User, error) {
	rows, err := dbpool.Query(ctx, "SELECT id_user, email, username FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.UserId, &user.Email, &user.Username)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (ur *UserRepositoryImpl) FindById(ctx context.Context, dbpool *pgxpool.Pool, id int) (domain.User, error) {
	var user domain.User
	err := dbpool.QueryRow(ctx, "SELECT id_user, email, username FROM users WHERE id_user = $1", id).Scan(&user.UserId, &user.Email, &user.Username)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (ur *UserRepositoryImpl) Create(ctx context.Context, dbpool *pgxpool.Pool, user domain.User) (domain.User, error) {
	err := dbpool.QueryRow(ctx, "INSERT INTO users (email, username) VALUES ($1, $2) RETURNING id_user", user.Email, user.Username).Scan(&user.UserId)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (ur *UserRepositoryImpl) Update(ctx context.Context, dbpool *pgxpool.Pool, user domain.User) (domain.User, error) {
	_, err := dbpool.Exec(ctx, "UPDATE users SET email = $1, username = $2 WHERE id_user = $3", user.Email, user.Username, user.UserId)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (ur *UserRepositoryImpl) Delete(ctx context.Context, dbpool *pgxpool.Pool, id int) error {
	_, err := dbpool.Exec(ctx, "DELETE FROM users WHERE id_user = $1", id)
	if err != nil {
		return err
	}

	return nil
}


