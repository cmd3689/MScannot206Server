package login

import (
	"MScannot206/pkg/login/mongodb"
	"MScannot206/shared/repository"
	"MScannot206/shared/service"
	"encoding/json"
	"errors"
	"net/http"
)

func NewLoginService(
	host service.ServiceHost,
	router *http.ServeMux,
) (*LoginService, error) {
	if host == nil {
		return nil, errors.New("host is null")
	}

	if router == nil {
		return nil, errors.New("router is null")
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
}

func (s *LoginService) Init(host service.ServiceHost) error {
	var err error

	// TODO: 외부에서 가져올 수 있도록 수정 필요
	dbName := "MStest"

	s.userRepo, err = mongodb.NewUserRepository(s.host.GetContext(), s.host.GetMongoClient(), dbName)
	if err != nil {
		return err
	}

	s.router.HandleFunc("/login", s.onLogin)

	return nil
}

func (s *LoginService) Start() error {
	var errs error

	if err := s.userRepo.Start(); err != nil {
		errs = errors.Join(errs, err)
	}

	if errs != nil {
		return errs
	}

	return nil
}

func (s *LoginService) Stop() error {
	var errs error

	if err := s.userRepo.Stop(); err != nil {
		errs = errors.Join(errs, err)
	}

	if errs != nil {
		return errs
	}

	return nil
}

func (s *LoginService) onLogin(w http.ResponseWriter, r *http.Request) {

	// HTTP POST만 허용
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// uid := req.Uid

	// user, err := s.userRepo.FindUserByUID(uid)
	// if err != nil {
	// 	http.Error(w, "User Not Found", http.StatusInternalServerError)
	// 	return
	// }

	// _ = user

	// TODO: 접속중인 계정인지 확인
	//println("User logged in:", req.UserIds)
	for _, id := range req.UserIds {
		println(id)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
