package service

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/masrayfa/go-delos-aqua/internals/models/domain"
	"github.com/masrayfa/go-delos-aqua/internals/models/web"
	"github.com/masrayfa/go-delos-aqua/internals/repository"
)

type UserServiceImpl struct {
	userRepository repository.UserRepository
	db             *pgxpool.Pool
}

func NewUserService(userRepository repository.UserRepository, db *pgxpool.Pool) UserService {
	return &UserServiceImpl{
		userRepository: userRepository,
		db:             db,
	}
}

func (us *UserServiceImpl) FindAll(ctx context.Context) ([]web.UserRead, error) {
	users, err := us.userRepository.FindAll(ctx, us.db)
	if err != nil {
		return nil, err
	}

	var userReads []web.UserRead
	for _, user := range users {
		userReads = append(userReads, web.UserRead{
			UserId:       user.UserId,
			Username: user.Username,
			Email:    user.Email,
		})
	}

	return userReads, nil
}

func (us *UserServiceImpl) FindById(ctx context.Context, id int) (web.UserRead, error) {
	user, err := us.userRepository.FindById(ctx, us.db, id)
	if err != nil {
		return web.UserRead{}, err
	}

	userRead := web.UserRead{
		UserId:       user.UserId,
		Username: user.Username,
		Email:    user.Email,
	}

	return userRead, nil
}

func (us *UserServiceImpl) Create(ctx context.Context, payload web.UserCreate) (web.UserRead, error) {
	userDomain := domain.User{
		Username: payload.Username,
		Email:    payload.Email,
		Password: payload.Password,
	}

	user, err := us.userRepository.Create(ctx, us.db, userDomain)
	if err != nil {
		return web.UserRead{}, err
	}

	userRead := web.UserRead{
		Username: user.Username,
		Email:    user.Email,
	}

	return userRead, nil
}

func (us *UserServiceImpl) Update(ctx context.Context, payload web.UserUpdate) (web.UserRead, error) {
	userDomain := domain.User{
		Username: payload.Username,
		Email:    payload.Email,
	}

	user, err := us.userRepository.Update(ctx, us.db, userDomain)
	if err != nil {
		return web.UserRead{}, err
	}

	userRead := web.UserRead{
		Username: user.Username,
		Email:    user.Email,
	}

	return userRead, nil
}

func (us *UserServiceImpl) Delete(ctx context.Context, id int) error {
	err := us.userRepository.Delete(ctx, us.db, id)
	if err != nil {
		return err
	}

	return nil
}
