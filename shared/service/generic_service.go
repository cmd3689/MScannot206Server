package service

type GenericService interface {
	Init(host ServiceHost) error
	Start() error
	Stop() error
}
