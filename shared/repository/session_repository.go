package repository

import (
	"MScannot206/shared/entity"
	"context"
)

type SessionRepository interface {
	SaveUserSessions(ctx context.Context, sessions []*entity.UserSession) error
}
