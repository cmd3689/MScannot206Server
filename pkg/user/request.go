package user

type UserNameCheckInfo struct {
	Uid   string `json:"uid"`
	Token string `json:"token"`
	Name  string `json:"name"`
}

type CheckCharacterNameRequest struct {
	Requests []*UserNameCheckInfo `json:"requests"`
}

type UserCreateCharacterInfo struct {
	Uid   string `json:"uid"`
	Token string `json:"token"`
	Slot  int    `json:"slot"`
	Name  string `json:"name"`
}

type CreateCharacterRequest struct {
	Requests []*UserCreateCharacterInfo `json:"requests"`
}

type UserDeleteCharacterInfo struct {
	Uid   string `json:"uid"`
	Token string `json:"token"`
	Slot  int    `json:"slot"`
	Name  string // not used in request, for internal use only
}

type DeleteCharacterRequest struct {
	Requests []*UserDeleteCharacterInfo `json:"requests"`
}
