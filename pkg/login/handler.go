package login

import (
	"MScannot206/shared/entity"
	"context"
	"errors"
)

var ErrAuthServiceHandlerIsNil = errors.New("auth service handler is null")

type AuthServiceHandler interface {
	CreateUserSessions(ctx context.Context, user []*entity.User) ([]*entity.UserSession, []*entity.User, error)
}

var ErrUserRepositoryHandlerIsNil = errors.New("user repository handler is null")

type UserRepositoryHandler interface {
	FindUserByUids(ctx context.Context, uids []string) ([]*entity.User, []string, error)
	InsertUserByUids(ctx context.Context, uids []string) ([]*entity.User, []string, error)
}
