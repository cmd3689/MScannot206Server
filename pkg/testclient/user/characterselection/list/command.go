package list

import (
	"MScannot206/pkg/testclient/framework"
	"MScannot206/pkg/testclient/user/handler"
	"sort"

	"github.com/rs/zerolog/log"
)

func NewCharacterListCommand(client framework.Client, userHandler handler.UserHandler) (*CharacterListCommand, error) {
	if client == nil {
		return nil, framework.ErrClientIsNil
	}

	if userHandler == nil {
		return nil, handler.ErrUserHandlerIsNil
	}

	return &CharacterListCommand{
		client:      client,
		userHandler: userHandler,
	}, nil
}

type CharacterListCommand struct {
	client      framework.Client
	userHandler handler.UserHandler
}

func (c *CharacterListCommand) Commands() []string {
	return []string{"character_list"}
}

func (c *CharacterListCommand) Execute(args []string) error {
	characterHandlers := c.userHandler.GetCharacterHandlers()

	// Key 기준으로 오름차순 정렬
	sort.Slice(characterHandlers, func(i, j int) bool {
		return characterHandlers[i].GetKey() < characterHandlers[j].GetKey()
	})

	log.Info().Msg("=== Character List ===")
	for _, ch := range characterHandlers {
		log.Info().Int("Key", ch.GetKey()).Str("Name", ch.GetName()).Msg("")
	}
	log.Info().Msg("======================")

	return nil
}

func (c *CharacterListCommand) Description() string {
	return framework.MakeCommandDescription(c.Commands(), "", "캐릭터 리스트를 요청 합니다.")
}
