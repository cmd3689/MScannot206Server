package handler

import (
	"MScannot206/pkg/testclient/framework"
	"errors"
)

var ErrUserHandlerIsNil = errors.New("user handler is nil")

type UserHandler interface {
	GetUid() string
	GetToken() string
	GetCharacterHandler(int) (CharacterHandler, bool)
	GetCharacterHandlers() []CharacterHandler

	AddCommand(cmd framework.ClientCommand) error
}

type CharacterHandler interface {
	GetKey() int
	GetName() string
}
