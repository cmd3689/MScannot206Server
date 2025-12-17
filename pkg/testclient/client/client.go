package client

import (
	"MScannot206/pkg/testclient/config"
	"MScannot206/pkg/testclient/framework"
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

const HTTP_TIMEOUT = 30 * time.Second

func NewClient(ctx context.Context, cfg *config.ClientConfig) (*Client, error) {
	if ctx == nil {
		return nil, errors.New("context is nil")
	}

	webClientCfg := cfg
	if webClientCfg == nil {
		webClientCfg = &config.ClientConfig{
			Url:  "http://localhost",
			Port: 8080,
		}
	}

	client := http.Client{
		Timeout: HTTP_TIMEOUT,
	}

	url := webClientCfg.Url + ":" + fmt.Sprintf("%v", webClientCfg.Port)
	if url == "" {
		return nil, errors.New("웹 클라이언트 URL이 비어있습니다")
	}

	ctxWithCancel, cancel := context.WithCancel(ctx)

	self := &Client{
		ctx:        ctxWithCancel,
		cancelFunc: cancel,

		cfg: webClientCfg,
		url: url,

		client: &client,

		logics: make([]framework.Logic, 0, 8),
	}

	return self, nil
}

type Client struct {
	framework.InputMachine

	ctx        context.Context
	cancelFunc context.CancelFunc

	// Config
	cfg *config.ClientConfig
	url string

	client *http.Client

	logics []framework.Logic
}

func (c *Client) GetContext() context.Context {
	return c.ctx
}

func (c *Client) Init() error {
	var errs error

	c.InputMachine.Init()

	for _, l := range c.logics {
		if err := l.Init(); err != nil {
			errs = errors.Join(errs, err)
			log.Err(err)
		}
	}

	if errs != nil {
		return errs
	}

	return nil
}

func (c *Client) Start() error {
	for _, l := range c.logics {
		if err := l.Start(); err != nil {
			return err
		}
	}

	c.InputMachine.Attach(c.ctx, c)

	<-c.ctx.Done()
	return nil
}

func (c *Client) Quit() error {
	for _, l := range c.logics {
		if err := l.Stop(); err != nil {
			return err
		}
	}

	c.InputMachine.Detach()

	if c.cancelFunc != nil {
		c.cancelFunc()
	}

	return nil
}

func (c *Client) AddLogic(logic framework.Logic) error {
	if logic == nil {
		return errors.New("logic is nil")
	}

	c.logics = append(c.logics, logic)
	return nil
}

func (c *Client) GetLogics() []framework.Logic {
	return c.logics
}

func (c *Client) AddCommand(cmd framework.ClientCommand) error {
	return c.InputMachine.AddCommand(cmd)
}

func (c *Client) GetUrl() string {
	return c.url
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.client.Do(req)
}
