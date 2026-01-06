package user

type UserEntity struct {
	Uid   string
	Token string
}

type UserCreateCharacter struct {
	Uid  string
	Slot int
	Name string
}

type UserDeleteCharacter struct {
	Uid  string
	Slot int
	Name string
}
