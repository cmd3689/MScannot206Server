package service

import (
	"MScannot206/shared/def"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

var ErrServiceHostIsNil = fmt.Errorf("service host is null")

type ServiceHost interface {
	// Core
	GetContext() context.Context
	GetServices() []Service
	AddService(svc Service) error
	Quit() error

	// DB
	GetMongoClient() *mongo.Client

	// ETC
	GetLocale() def.Locale
}

func GetService[T Service](host ServiceHost) (T, error) {
	var ret T

	if host == nil {
		return ret, fmt.Errorf("host is nil")
	}

	for _, svc := range host.GetServices() {
		if casted, ok := svc.(T); ok {
			return casted, nil
		}
	}

	return ret, fmt.Errorf("service not found: %T", ret)
}
