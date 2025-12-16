package service

type GenericService interface {
	GetPriority() int
	Init() error
	Start() error
	Stop() error
}
