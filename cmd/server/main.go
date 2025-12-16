package main

import (
	"MScannot206/pkg/logger"
	"MScannot206/shared/config"
	"context"
	"time"

	"github.com/rs/zerolog/log"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Error().Msgf("프로그램 비정상 종료: %v", r)
		}
	}()

	logCfg := &config.LogConfig{
		AppName:   "server",
		DebugMode: true,
	}

	serverCfg := &config.WebServerConfig{
		Port: 8080,

		ServerName: "DevServer",

		MongoUri:       "mongodb://localhost:27017/",
		MongoEnvDBName: "MSenv",
	}

	if err := setupConfig(logCfg, serverCfg); err != nil {
		log.Err(err)
		panic(err)
	}

	if err := logger.GetLogManager().Init(*logCfg); err != nil {
		log.Err(err)
		panic(err)
	}
	defer logger.GetLogManager().Close()

	for {
		log.Info().Msg("서버 초기화를 시작합니다.")

		err := run(context.Background(), serverCfg)
		if err == nil {
			log.Info().Msg("서버가 정상적으로 종료되었습니다.")
			break
		} else {
			log.Error().Err(err).Msg("서버가 비정상적으로 종료되었습니다. 3초 후 재시작합니다.")
			time.Sleep(3 * time.Second)
		}
	}

}
