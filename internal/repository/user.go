package repository

import (
	"context"
	"errors"
	"gotest/internal/model"

	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user model.UserModel) (*uuid.UUID, error)
	FindUser(ctx context.Context, userId uuid.UUID) (*model.UserModel, error)
	DeleteUser(ctx context.Context, userId uuid.UUID) error
	UpdateUser(ctx context.Context, userId uuid.UUID, user model.UserModel) error
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}


func (r *userRepository) CreateUser(ctx context.Context, user model.UserModel) (*uuid.UUID, error) {
	return nil, errors.New("aaaa")
}

func (r *userRepository) FindUser(ctx context.Context, userId uuid.UUID) (*model.UserModel, error) {
	return nil, errors.New("aaaa")
}

func (r *userRepository) DeleteUser(ctx context.Context, userId uuid.UUID) error {
	return errors.New("aaaa")
}

func (r *userRepository) UpdateUser(ctx context.Context, userId uuid.UUID, user model.UserModel) error {
	return errors.New("aaaa")
}
