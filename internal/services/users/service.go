package users

import (
	"context"
	usersC "github.com/Rasikrr/learning_platform_gateway/internal/clients/users"
	"github.com/Rasikrr/learning_platform_gateway/internal/domain/entity"
)

type Service interface {
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
}

type service struct {
	client usersC.Client
}

func NewService(
	client usersC.Client,
) Service {
	return &service{
		client: client,
	}
}

func (s *service) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	return s.client.GetByEmail(ctx, email)
}
