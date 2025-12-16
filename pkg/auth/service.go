package auth

import (
	"MScannot206/pkg/auth/session"
	"MScannot206/pkg/serverinfo"
	"MScannot206/shared/entity"
	"MScannot206/shared/repository"
	"MScannot206/shared/service"
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"

	"github.com/rs/zerolog/log"
)

func NewAuthService(
	host service.ServiceHost,
) (*AuthService, error) {
	if host == nil {
		return nil, errors.New("host is null")
	}

	return &AuthService{
		host: host,
	}, nil
}

type AuthService struct {
	host service.ServiceHost

	sessionRepo repository.SessionRepository
}

func (s *AuthService) GetPriority() int {
	return 0
}

func (s *AuthService) Init() error {
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

	s.sessionRepo, err = session.NewSessionRepository(s.host.GetContext(), s.host.GetMongoClient(), gameDBName)
	if err != nil {
		return err
	}

	return errs
}

func (s *AuthService) Start() error {
	return nil
}

func (s *AuthService) Stop() error {
	return nil
}

func generateToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func (s *AuthService) CreateUserSessions(ctx context.Context, user []*entity.User) ([]*entity.UserSession, []*entity.User, error) {
	sessions := make([]*entity.UserSession, 0, len(user))
	failureUsers := make([]*entity.User, 0)

	for _, u := range user {
		token, err := generateToken()
		if err != nil {
			log.Warn().Err(err)
			continue
		}

		session := &entity.UserSession{
			Uid:   u.Uid,
			Token: token,
		}

		sessions = append(sessions, session)
	}

	err := s.sessionRepo.SaveUserSessions(ctx, sessions)
	if err != nil {
		return nil, nil, err
	}

	return sessions, failureUsers, nil
}
