package main

import (
	"MScannot206/pkg/login"
	"MScannot206/shared/config"
	"MScannot206/shared/server"
	"context"
	"errors"
	"flag"
	"log"
)

func main() {
	var errs error

	var cfgPath string

	flag.StringVar(&cfgPath, "config", "", "서버 설정 파일 경로 지정")
	flag.Parse()

	var cfg *config.WebServerConfig
	if cfgPath != "" {
		if err := config.LoadYamlConfig(cfgPath, cfg); err != nil {
			log.Default().Printf("서버 설정 파일 로드 실패:%v", err)
		}
	}

	web_server, err := server.NewWebServer(
		context.Background(),
		cfg,
	)

	if err != nil {
		panic(err)
	}

	// 로그인 서비스
	if login_service, err := login.NewLoginService(web_server, web_server.GetRouter()); err != nil {
		errs = errors.Join(errs, err)
		log.Println(err)
	} else {
		if err := web_server.AddService(login_service); err != nil {
			errs = errors.Join(errs, err)
			log.Println(err)
		}
	}

	if errs != nil {
		panic(errs)
	}

	if err := web_server.Init(); err != nil {
		panic(err)
	}

	go func() {
		if err := web_server.Start(); err != nil {
			panic(err)
		}
	}()

	<-web_server.GetContext().Done()

	log.Println("웹 서버 종료")
}
