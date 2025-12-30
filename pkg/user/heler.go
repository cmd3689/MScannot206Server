package user

import (
	"MScannot206/shared/def"
	"MScannot206/shared/util"
	"unicode/utf8"
)

// 유저 캐릭터 슬롯 판별
func IsInvalidCharacterSlot(slot int) bool {
	return slot < 1 || slot > def.MaxCharacterSlot
}

// 유저 캐릭터 이름 유효성 검사
func ValidateCharacterName(name string, locale def.Locale) string {
	l := utf8.RuneCountInString(name)
	if l < def.MinCharacterNameLength {
		return USER_CREATE_CHARACTER_NAME_MIN_LENGTH_ERROR
	} else if l > def.MaxCharacterNameLength {
		return USER_CREATE_CHARACTER_NAME_MAX_LENGTH_ERROR
	} else if util.HasSpecialChar(name, locale) {
		return USER_CREATE_CHARACTER_NAME_SPECIAL_CHAR_ERROR
	}
	return ""
}
