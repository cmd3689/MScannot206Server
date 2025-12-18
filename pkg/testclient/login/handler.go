package login

import (
	"MScannot206/pkg/testclient/user"
	"MScannot206/shared/entity"
	"errors"
)

var ErrUserLogicHandlerIsNil = errors.New("user logic handler is nil")

type UserLogicHandler interface {
	ConnectUser(userEntity *entity.User, token string) (*user.User, error)
	DisconnectUser(uid string) error
}
