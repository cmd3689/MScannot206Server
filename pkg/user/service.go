package user

import (
	"MScannot206/pkg/serverinfo"
	"MScannot206/shared/repository"
	"MScannot206/shared/server"
	"MScannot206/shared/service"
	"errors"
	"net/http"

	"github.com/rs/zerolog/log"
)

func NewUserService(
	host service.ServiceHost,
	router *http.ServeMux,
) (*UserService, error) {
	if host == nil {
		return nil, service.ErrServiceHostIsNil
	}

	if router == nil {
		return nil, server.ErrServeMuxIsNil
	}

	return &UserService{
		host:   host,
		router: router,
	}, nil
}

type UserService struct {
	host   service.ServiceHost
	router *http.ServeMux

	userRepo repository.UserRepository
}

func (s *UserService) GetPriority() int {
	return 0
}

func (s *UserService) Init() error {
	var errs error
	var err error
	var gameDBName string = ""

	serverInfoService, err := service.GetService[*serverinfo.ServerInfoService](s.host)
	if err != nil {
		log.Err(err)
		errs = errors.Join(errs, err)
	} else {
		srvInfo, err := serverInfoService.GetInfo()
		if err != nil {
			log.Err(err)
			errs = errors.Join(errs, err)
		} else {
			gameDBName = srvInfo.GameDBName
		}
	}

	s.userRepo, err = NewUserRepository(s.host.GetMongoClient(), gameDBName)
	if err != nil {
		return err
	}

	return errs
}

func (s *UserService) Start() error {
	return nil
}

func (s *UserService) Stop() error {
	return nil
}
