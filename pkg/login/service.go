package login

import (
	"MScannot206/pkg/serverinfo"
	"MScannot206/pkg/user"
	"MScannot206/shared/entity"
	"MScannot206/shared/repository"
	"MScannot206/shared/server"
	"MScannot206/shared/service"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/rs/zerolog/log"
)

func NewLoginService(
	host service.ServiceHost,
	router *http.ServeMux,
) (*LoginService, error) {
	if host == nil {
		return nil, service.ErrServiceHostIsNil
	}

	if router == nil {
		return nil, server.ErrServeMuxIsNil
	}

	return &LoginService{
		host:   host,
		router: router,
	}, nil
}

type LoginService struct {
	host   service.ServiceHost
	router *http.ServeMux

	userRepo repository.UserRepository

	authServiceHandler AuthServiceHandler
}

func (s *LoginService) GetPriority() int {
	return 0
}

func (s *LoginService) Init() error {
	var errs error
	var err error
	var gameDBName string = ""

	s.router.HandleFunc("/login", s.onLogin)

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

	s.userRepo, err = user.NewUserRepository(s.host.GetMongoClient(), gameDBName)
	if err != nil {
		log.Err(err)
		errs = errors.Join(errs, err)
	}

	return errs
}

func (s *LoginService) Start() error {
	return nil
}

func (s *LoginService) Stop() error {
	return nil
}

func (s *LoginService) SetHandlers(
	authService AuthServiceHandler,
) error {
	var errs error

	s.authServiceHandler = authService
	if authService == nil {
		errs = errors.Join(errs, ErrAuthServiceHandlerIsNil)
	}

	return errs
}

func (s *LoginService) onLogin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var response LoginResponse

	users, newUids, err := s.userRepo.FindUserByUids(ctx, req.Uids)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 신규 유저 생성
	if len(newUids) > 0 {
		newUsers, err := s.userRepo.InsertUserByUids(ctx, newUids)
		if err != nil {
			// 신규 유저는 로그인 불가
			log.Printf("신규 유저 생성 불가: %v", err)

			for _, uid := range newUids {
				reason := LoginFailure{
					Uid:       uid,
					ErrorCode: LOGIN_DB_WRITE_ERROR,
				}
				response.FailUids = append(response.FailUids, reason)
			}

			// TODO: 생성 불가능한 유저 uid 로그 추가
		} else {
			users = append(users, newUsers...)
		}
	}

	sessions, failureUsers, err := s.authServiceHandler.CreateUserSessions(ctx, users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	usersByUid := make(map[string]*entity.User)
	for _, u := range users {
		usersByUid[u.Uid] = u
	}

	for _, session := range sessions {
		if u, ok := usersByUid[session.Uid]; ok {
			success := LoginSuccess{
				UserEntity: u,
				Token:      session.Token,
			}
			response.SuccessUids = append(response.SuccessUids, success)
		} else {
			log.Warn().Msgf("세션은 존재하나 유저가 없음: %s", session.Uid)
		}
	}

	for _, u := range failureUsers {
		reason := LoginFailure{
			Uid:       u.Uid,
			ErrorCode: LOGIN_SESSION_CREATE_ERROR,
		}
		response.FailUids = append(response.FailUids, reason)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Err(err)
	}
}
