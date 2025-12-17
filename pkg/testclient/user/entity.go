package user

import (
	"MScannot206/pkg/testclient/framework"
	"context"
	"errors"
)

type User struct {
	framework.InputMachine

	Uid   string
	Token string
}

func (u *User) Init() {
	u.InputMachine.Init()
}

func (u *User) Attach(ctx context.Context) error {
	if ctx == nil {
		return errors.New("ctx is null")
	}

	u.InputMachine.Attach(ctx, u)
	<-u.Done()
	return nil
}

func (u *User) Quit() error {
	u.InputMachine.Detach()
	return nil
}
