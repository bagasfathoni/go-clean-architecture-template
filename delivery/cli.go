package delivery

import "github.com/bagasfathoni/go-clean-architecture-template/usecases"

type aCli struct {
	repo usecases.AUsecases
}

type ACli interface {
	StartCli()
}

func (a *aCli) StartCli() {
	panic("Needs implementation")
}

func InitACli(a usecases.AUsecases) ACli {
	aNewCli := new(aCli)
	aNewCli.repo = a
	return aNewCli
}
