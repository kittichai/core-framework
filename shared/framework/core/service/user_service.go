package service

import (
	"context"

	repository "github.com/kittichai/core-framework/shared/framework/infrastructure/persistence/repository"
	model "github.com/kittichai/core-model/shared/model/core/domain"
	repo_model "github.com/kittichai/core-model/shared/model/infrastructure/persistence/model"
)

type UserService interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	FindByUsername(ctx context.Context, username string) (*model.User, error)
}

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserServiceImpl(userRepository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}

func (s *UserServiceImpl) Create(ctx context.Context, user *model.User) error {
	repo_model_user := &repo_model.User{
		ID:           user.ID,
		Email:        user.Email,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		UserName:     user.UserName,
		PasswordHash: user.PasswordHash,
		CreatedAt:    user.CreatedAt,
	}
	return s.userRepository.Create(ctx, repo_model_user)
}

func (s *UserServiceImpl) Update(ctx context.Context, user *model.User) error {
	repo_model_user := &repo_model.User{
		ID:           user.ID,
		Email:        user.Email,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		UserName:     user.UserName,
		PasswordHash: user.PasswordHash,
		CreatedAt:    user.CreatedAt,
	}
	return s.userRepository.Update(ctx, repo_model_user)
}

func (s *UserServiceImpl) FindByEmail(ctx context.Context, email string) (*model.User, error) {

	user_repo, err := s.userRepository.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		ID:        user_repo.ID,
		Email:     user_repo.Email,
		FirstName: user_repo.FirstName,
		LastName:  user_repo.LastName,
		UserName:  user_repo.UserName,
		CreatedAt: user_repo.CreatedAt,
	}

	return user, nil
}

func (s *UserServiceImpl) FindByUsername(ctx context.Context, username string) (*model.User, error) {

	user_repo, err := s.userRepository.FindByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		ID:        user_repo.ID,
		Email:     user_repo.Email,
		FirstName: user_repo.FirstName,
		LastName:  user_repo.LastName,
		UserName:  user_repo.UserName,
		CreatedAt: user_repo.CreatedAt,
	}

	return user, nil
}
