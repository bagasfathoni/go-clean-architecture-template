package usecases

import "github.com/bagasfathoni/go-clean-architecture-template/repository"

type aUsecases struct {
	repo repository.ARepository
}

type AUsecases interface {
	FooForBar()
}

func (a *aUsecases) FooForBar() {
	panic("Needs implemantion")
}

func InitCustomerUsecases(a repository.ARepository) AUsecases {
	aUsec := new(aUsecases)
	aUsec.repo = a
	return aUsec
}
