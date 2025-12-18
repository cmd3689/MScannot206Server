package login

import (
	"MScannot206/pkg/testclient/framework"
)

func NewLoginCommand(client framework.Client) (*LoginCommand, error) {
	if client == nil {
		return nil, framework.ErrClientIsNil
	}

	loginLogic, err := framework.GetLogic[*LoginLogic](client)
	if err != nil {
		return nil, err
	}

	return &LoginCommand{
		client: client,

		loginLogic: loginLogic,
	}, nil
}

type LoginCommand struct {
	client framework.Client

	loginLogic *LoginLogic
}

func (c *LoginCommand) Commands() []string {
	return []string{"login"}
}

func (c *LoginCommand) Execute(args []string) error {
	if len(args) < 1 {
		return framework.ErrInvalidCommandArgument
	}

	uid := args[0]

	if err := c.loginLogic.RequestLogin(uid); err != nil {
		return err
	}

	return nil
}

func (c *LoginCommand) Description() string {
	return framework.MakeCommandDescription(c.Commands(), "<uid>", "로그인을 요청 합니다.")
}
