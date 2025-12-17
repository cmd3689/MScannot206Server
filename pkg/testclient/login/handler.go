package login

import "errors"

var UserLogicHandlerIsNil = errors.New("user logic handler is nil")

type UserLogicHandler interface {
	ConnectUser(uid string, token string) error
	DisconnectUser(uid string) error
}
