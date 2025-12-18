package user

import (
	"MScannot206/pkg/testclient/framework"
	"MScannot206/pkg/testclient/user/character"
	"MScannot206/pkg/testclient/user/handler"
	"context"
	"errors"

	"github.com/rs/zerolog/log"
)

var ErrUserIsNil = errors.New("user is nil")

func NewUser(uid string, token string) (*User, error) {
	if uid == "" {
		return nil, ErrUserIsNil
	}

	u := &User{
		Uid:   uid,
		Token: token,

		Characters: []*character.Character{},
	}

	u.InputMachine.Init()

	return u, nil
}

type User struct {
	framework.InputMachine

	Uid   string
	Token string

	Characters []*character.Character
}

func (u *User) GetUid() string {
	return u.Uid
}

func (u *User) GetToken() string {
	return u.Token
}

func (u *User) GetCharacterHandler(key int) (handler.CharacterHandler, bool) {
	for _, ch := range u.Characters {
		if ch.GetKey() == key {
			return ch, true
		}
	}
	return nil, false
}

func (u *User) GetCharacterHandlers() []handler.CharacterHandler {
	ret := make([]handler.CharacterHandler, len(u.Characters))
	for i, ch := range u.Characters {
		ret[i] = ch
	}
	return ret
}

func (u *User) Attach(ctx context.Context) error {
	if ctx == nil {
		return errors.New("ctx is null")
	}

	u.InputMachine.Attach(ctx, u)

	log.Info().Msgf("유저 %s를 선택하였습니다.", u.Uid)

	<-u.Done()
	return nil
}

func (u *User) Quit() error {
	log.Info().Msgf("유저 %s를 선택 해제하였습니다.", u.Uid)
	u.InputMachine.Detach()
	return nil
}
