package repository

import (
	"MScannot206/shared/entity"
	"context"
)

type UserRepository interface {
	FindUserByUids(ctx context.Context, uids []string) ([]*entity.User, []string, error)
	InsertUserByUids(ctx context.Context, uids []string) ([]*entity.User, error)
}
