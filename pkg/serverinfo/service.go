package serverinfo

import (
	"MScannot206/shared/service"
	"errors"
)

func NewServerInfoService(host service.ServiceHost, serverName, dbName string) (*ServerInfoService, error) {
	if host == nil {
		return nil, errors.New("service host is nil")
	}

	return &ServerInfoService{
		host:       host,
		serverName: serverName,
		dbName:     dbName,
	}, nil
}

type ServerInfoService struct {
	host       service.ServiceHost
	serverName string
	dbName     string

	serverInfoRepo *ServerInfoRepository
}

func (s *ServerInfoService) GetPriority() int {
	return 0
}

func (s *ServerInfoService) Init() error {
	var err error

	s.serverInfoRepo, err = NewServerInfoRepository(s.host.GetMongoClient(), s.dbName)
	if err != nil {
		return err
	}

	info, err := s.serverInfoRepo.GetInfo(s.host.GetContext(), s.serverName)
	if err != nil {
		return err
	} else if info == nil {
		info = &ServerInfo{
			Name: s.serverName,

			GameDBName: "MSgame",
			LogDBName:  "MSlog",

			Status: StatusActive,

			Description: "자동 생성 된 서버 정보",
		}

		if err := s.serverInfoRepo.SetInfo(s.host.GetContext(), info); err != nil {
			return err
		}
	}

	return nil
}

func (s *ServerInfoService) Start() error {
	return nil
}

func (s *ServerInfoService) Stop() error {
	return nil
}

func (s *ServerInfoService) GetInfo() (*ServerInfo, error) {
	return s.serverInfoRepo.GetInfo(s.host.GetContext(), s.serverName)
}
