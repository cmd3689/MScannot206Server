package command

import (
	"MScannot206/pkg/testclient/framework"
	"MScannot206/pkg/testclient/user/characterselection/create"
	"MScannot206/pkg/testclient/user/characterselection/list"
	"MScannot206/pkg/testclient/user/handler"
	"errors"

	"github.com/rs/zerolog/log"
)

func RegisterCommands(client framework.Client, userHandler handler.UserHandler) error {
	var errs error

	if client == nil {
		return framework.ErrClientIsNil
	}

	if userHandler == nil {
		return handler.ErrUserHandlerIsNil
	}

	characterListCmd, err := list.NewCharacterListCommand(client, userHandler)
	if err != nil {
		errs = errors.Join(errs, err)
		log.Err(err)
	}

	characterCreateCmd, err := create.NewCharacterCreateCommand(client, userHandler)
	if err != nil {
		errs = errors.Join(errs, err)
		log.Err(err)
	}

	if errs != nil {
		return errs
	}

	errs = nil
	for _, cmd := range []framework.ClientCommand{
		characterListCmd,
		characterCreateCmd,
	} {
		if err := userHandler.AddCommand(cmd); err != nil {
			return err
		}
	}

	return errs
}
