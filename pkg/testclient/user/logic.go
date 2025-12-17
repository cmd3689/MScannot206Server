package user

import "MScannot206/pkg/testclient/framework"

const userCapacity = 1000

func NewUserLogic(client framework.Client) (*UserLogic, error) {
	if client == nil {
		return nil, framework.ErrClientIsNil
	}
	return &UserLogic{
		client: client,

		users: make(map[string]*User, userCapacity),
	}, nil
}

type UserLogic struct {
	client framework.Client

	users map[string]*User
}

func (l *UserLogic) Init() error {
	return nil
}

func (l *UserLogic) Start() error {
	return nil
}

func (l *UserLogic) Stop() error {
	return nil
}

func (l *UserLogic) ConnectUser(uid string, token string) error {
	l.users[uid] = &User{
		Uid:   uid,
		Token: token,
	}

	l.users[uid].Init()
	return nil
}

func (l *UserLogic) DisconnectUser(uid string) error {
	if user, ok := l.users[uid]; ok {
		user.Quit()
		delete(l.users, uid)
	}
	return nil
}

func (l *UserLogic) GetUser(uid string) (*User, bool) {
	user, ok := l.users[uid]
	return user, ok
}
