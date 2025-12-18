package userselection

import (
	"MScannot206/pkg/testclient/framework"
	"MScannot206/pkg/testclient/user"
	"fmt"
)

func NewUserSelectionCommand(client framework.Client) (*UserSelectionCommand, error) {
	if client == nil {
		return nil, framework.ErrClientIsNil
	}

	return &UserSelectionCommand{
		client: client,
	}, nil
}

type UserSelectionCommand struct {
	client framework.Client
}

func (c *UserSelectionCommand) Commands() []string {
	return []string{"user_select"}
}

func (c *UserSelectionCommand) Execute(args []string) error {
	userLogic, err := framework.GetLogic[*user.UserLogic](c.client)
	if err != nil {
		return err
	}

	if len(args) < 1 {
		return framework.ErrInvalidCommandArgument
	}

	uid := args[0]
	user, ok := userLogic.GetUser(uid)
	if !ok {
		return fmt.Errorf("유저를 찾을 수 없습니다: %s", uid)
	}

	return user.Attach(c.client.GetContext())
}

func (c *UserSelectionCommand) Description() string {
	return framework.MakeCommandDescription(c.Commands(), "<uid>", "유저 선택을 요청 합니다.")
}
