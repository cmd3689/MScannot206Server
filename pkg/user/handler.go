package user

import (
	"MScannot206/shared/entity"
	"context"
	"errors"
)

var ErrServerInfoServiceHandlerIsNil = errors.New("server info service handler is null")

type ServerInfoServiceHandler interface {
	GetGameDBName() (string, error)
	GetLogDBName() (string, error)
}

var ErrAuthServiceHandlerIsNil = errors.New("auth service handler is null")

type AuthServiceHandler interface {
	ValidateUserSessions(ctx context.Context, sessions []*entity.UserSession) ([]string, []string, error)
}
