package character

import "errors"

var ErrCharacterIsNil = errors.New("character is null")
var ErrCharacterKeyIsInvalid = errors.New("캐릭터 키가 유효하지 않습니다")
var ErrCharacterNameIsEmpty = errors.New("캐릭터 이름이 비어있습니다")

func NewCharacter(slot int, name string) (*Character, error) {
	if name == "" {
		return nil, ErrCharacterNameIsEmpty
	}

	return &Character{
		Slot: slot,
		Name: name,
	}, nil
}

type Character struct {
	Slot int
	Name string
}

func (c Character) GetKey() int {
	return c.Slot
}

func (c Character) GetName() string {
	return c.Name
}
