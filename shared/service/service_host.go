package service

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceHost interface {
	// Core
	GetContext() context.Context

	// DB
	GetMongoClient() *mongo.Client
}

func GetService[T GenericService](host ServiceHost) T {
	var ret T
	// TODO: 서비스 가져오기
	return ret
}
