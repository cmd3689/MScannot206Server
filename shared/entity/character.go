package entity

import (
	"errors"
	"time"
)

var ErrCharacterSlotIsFull = errors.New("캐릭터 슬롯이 가득 찼습니다")

func NewCharacter(slot int, name string) *Character {
	return &Character{
		Slot: slot,
		Name: name,
	}
}

type Character struct {
	Slot int    `json:"slot" bson:"slot"`
	Name string `json:"name" bson:"name"`
}

var ErrCharacterNameIsNil = errors.New("character name entity is null")

type CharacterName struct {
	Name      string    `bson:"_id"`
	CreatedAt time.Time `bson:"create_at"`
}
