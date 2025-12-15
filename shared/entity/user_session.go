package entity

import "time"

type UserSession struct {
	Uid       string    `bson:"_id"`
	Token     string    `bson:"access_token"`
	UpdatedAt time.Time `bson:"updated_at"`
}
