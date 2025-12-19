package entity

import (
	"errors"
	"time"
)

var ErrCharacterSlotIsFull = errors.New("캐릭터 슬롯이 가득 찼습니다")

func NewCharacter(key int, name string) *Character {
	return &Character{
		Key:  key,
		Name: name,
	}
}

type Character struct {
	Key  int    `json:"key" bson:"key"`
	Name string `json:"name" bson:"name"`
}

var ErrCharacterNameIsNil = errors.New("character name entity is null")

type CharacterName struct {
	Name      string    `bson:"_id"`
	CreatedAt time.Time `bson:"create_at"`
}
