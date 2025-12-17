package framework

import (
	"context"
	"errors"
	"net/http"
)

var ErrClientIsNil = errors.New("client is nil")

type Client interface {
	GetContext() context.Context
	Init() error
	Start() error
	Quit() error

	AddLogic(logic Logic) error
	GetLogics() []Logic

	AddCommand(cmd ClientCommand) error

	// http
	GetUrl() string
	Do(req *http.Request) (*http.Response, error)
}
