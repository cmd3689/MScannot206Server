package characterselection

import (
	"MScannot206/pkg/testclient/framework"
)

func NewCharacterSelectionCommand(client framework.Client) (*CharacterSelectionCommand, error) {
	if client == nil {
		return nil, framework.ErrClientIsNil
	}

	return &CharacterSelectionCommand{
		client: client,
	}, nil
}

type CharacterSelectionCommand struct {
	client framework.Client
}

func (c *CharacterSelectionCommand) Commands() []string {
	return []string{"character_list"}
}

func (c *CharacterSelectionCommand) Execute(args []string) error {
	// userLogic, err := framework.GetLogic[*user.UserLogic](c.client)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (c *CharacterSelectionCommand) Description() string {
	return framework.MakeCommandDescription(c.Commands(), "", "캐릭터 리스트를 요청 합니다.")
}
