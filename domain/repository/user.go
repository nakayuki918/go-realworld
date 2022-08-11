package repository

import (
	"context"

	"github.com/nakayuki918/go-realworld/domain/entity"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) (entity.User, error)
	GetUser(ctx context.Context, userId int) (entity.User, error)
}



