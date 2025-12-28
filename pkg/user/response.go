package user

type UserNameCheckResult struct {
	Uid       string `json:"uid"`
	Available bool   `json:"available"`
	ErrorCode string `json:"error_code,omitempty"`
}

type CheckCharacterNameResponse struct {
	Responses []*UserNameCheckResult `json:"responses"`
}

type UserCreateCharacterResult struct {
	Uid       string `json:"uid"`
	Slot      int    `json:"slot,omitempty"`
	ErrorCode string `json:"error_code,omitempty"`
}

type CreateCharacterResponse struct {
	Responses []*UserCreateCharacterResult `json:"responses"`
}

type UserDeleteCharacterResult struct {
	Uid       string `json:"uid"`
	ErrorCode string `json:"error_code,omitempty"`
}

type DeleteCharacterResponse struct {
	Responses []*UserDeleteCharacterResult `json:"responses"`
}
