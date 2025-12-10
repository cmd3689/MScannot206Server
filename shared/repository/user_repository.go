package repository

func NewUser(uid string) *User {
	return &User{
		Uid: uid,
	}
}

type User struct {
	Uid string `json:"uid" bson:"uid"`
}

type UserRepository interface {
	Start() error
	Stop() error

	FindUserByUID(string) (*User, error)
}
