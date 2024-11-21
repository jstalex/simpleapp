package service

import (
	"context"

	"simpleapp/internal/domain"
	"simpleapp/internal/storage"
)

type Service interface {
	AddUser(ctx context.Context, user domain.User) error
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
}

type service struct {
	userStorage storage.UserStorage
}

func NewService(us storage.UserStorage) Service {
	return &service{
		userStorage: us,
	}
}

func (s *service) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	return s.userStorage.FindUser(ctx, email)
}

func (s *service) AddUser(ctx context.Context, user domain.User) error {
	return s.userStorage.InsertUser(ctx, user)
}
