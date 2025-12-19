package session

import "MScannot206/shared"

const SESSION_TOKEN_INVALID_ERROR = "SESSION_TOKEN_INVALID_ERROR"

func init() {
	shared.RegisterError(SESSION_TOKEN_INVALID_ERROR, "세션 토큰이 유효하지 않습니다.")
}
