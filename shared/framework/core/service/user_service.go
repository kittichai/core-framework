package service

import (
	"context"

	model "github.com/kittichai/core-model/shared/model/core/domain"
)

type UserService interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	FindByUsername(ctx context.Context, username string) (*model.User, error)
}

type UserServiceImpl struct {
	UserRepository UserService
}

func (s *UserServiceImpl) Create(ctx context.Context, user *model.User) error {
	return s.UserRepository.Create(ctx, user)
}

func (s *UserServiceImpl) Update(ctx context.Context, user *model.User) error {
	return s.UserRepository.Update(ctx, user)
}

func (s *UserServiceImpl) FindByEmail(ctx context.Context, email string) (*model.User, error) {

	user_repo, err := s.UserRepository.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		ID:    user_repo.ID,
		Email: user_repo.Email,
		Name:  user_repo.Name,
	}

	return user, nil
}

func (s *UserServiceImpl) FindByUsername(ctx context.Context, username string) (*model.User, error) {

	user_repo, err := s.UserRepository.FindByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		ID:    user_repo.ID,
		Email: user_repo.Email,
		Name:  user_repo.Name,
	}

	return user, nil
}
