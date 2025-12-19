package login

import "MScannot206/shared/entity"

type LoginSuccess struct {
	UserEntity *entity.User `json:"user_entity"`
	Token      string       `json:"token"`
}

type LoginFailure struct {
	Uid       string `json:"uid"`
	ErrorCode string `json:"error_code,omitempty"`
}

type LoginResponse struct {
	SuccessUids []*LoginSuccess `json:"success_uids"`
	FailUids    []*LoginFailure `json:"fail_uids"`
}
