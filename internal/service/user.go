package service

import (
	"context"
	"errors"
	"gotest/internal/model"
	"gotest/internal/repository"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(ctx context.Context, user model.UserModel) (*uuid.UUID, error)
	FindUser(ctx context.Context, userId uuid.UUID) (*model.UserModel, error)
	DeleteUser(ctx context.Context, userId uuid.UUID) error
	UpdateUser(ctx context.Context, userId uuid.UUID, user model.UserModel) error
}

type userService struct{
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}


func (s *userService) CreateUser(ctx context.Context, user model.UserModel) (*uuid.UUID, error) {
	return nil, errors.New("aaaa")
}

func (s *userService) FindUser(ctx context.Context, userId uuid.UUID) (*model.UserModel, error) {
	return nil, errors.New("aaaa")
}

func (s *userService) DeleteUser(ctx context.Context, userId uuid.UUID) error {
	return errors.New("aaaa")
}

func (s *userService) UpdateUser(ctx context.Context, userId uuid.UUID, user model.UserModel) error {
	return errors.New("aaaa")
}