package main

import (
	"MScannot206/shared/client"
	"MScannot206/shared/config"
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var cfgPath string

	flag.StringVar(&cfgPath, "config", "", "클라이언트 설정 파일 경로 지정")
	flag.Parse()

	var cfg *config.WebClientConfig
	if cfgPath != "" {
		if err := config.LoadYamlConfig(cfgPath, cfg); err != nil {
			log.Default().Printf("클라이언트 설정 파일 로드 실패:%v", err)
		}
	}

	client, err := client.NewWebClient(
		context.Background(),
		cfg,
	)

	if err != nil {
		panic(err)
	}

	if err := client.Init(); err != nil {
		panic(err)
	}

	go func() {
		if err := client.Start(); err != nil {
			panic(err)
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-sigCh:
		log.Printf("애플리케이션 강제 종료")

	case <-client.GetContext().Done():
		log.Printf("클라이언트 종료")
	}
}
