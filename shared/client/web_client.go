package client

import (
	"MScannot206/shared/config"
	"bufio"
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

const HTTP_TIMEOUT = 30 * time.Second

func NewWebClient(ctx context.Context, cfg *config.WebClientConfig) (*WebClient, error) {
	ctxWithCancel, cancel := context.WithCancel(ctx)

	client := http.Client{
		Timeout: HTTP_TIMEOUT,
	}

	webClientCfg := cfg
	if webClientCfg == nil {
		webClientCfg = &config.WebClientConfig{
			Url:  "http://localhost",
			Port: 8080,
		}
	}

	self := &WebClient{
		ctx:        ctxWithCancel,
		cancelFunc: cancel,

		client: &client,

		cfg: webClientCfg,

		inputChan: make(chan string),
	}

	return self, nil
}

type WebClient struct {
	ctx        context.Context
	cancelFunc context.CancelFunc

	// Config
	cfg *config.WebClientConfig

	// Core
	client *http.Client

	inputChan chan string
}

func (c *WebClient) GetContext() context.Context {
	return c.ctx
}

func (c *WebClient) Init() error {
	return nil
}

func (c *WebClient) Start() error {

	go c.taskClient()

	<-c.ctx.Done()

	return nil
}

func (c *WebClient) Shutdown() error {

	if c.cancelFunc != nil {
		c.cancelFunc()
	}

	return nil
}

func (c *WebClient) taskClient() {
	go c.inputClient()

	for {
		select {
		case input := <-c.inputChan:
			c.handleInput(input)

		case <-c.ctx.Done():
			return
		}
	}
}

func (c *WebClient) inputClient() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("입력 오류: %v\n", err)
			continue
		}

		tInput := strings.TrimSpace(input)
		if tInput == "" {
			continue
		}

		c.inputChan <- tInput
	}
}

func (c *WebClient) handleInput(input string) {
}
