package repository

import "github.com/bagasfathoni/go-clean-architecture-template/model"

type aRepository struct {
	repo model.A
}

type ARepository interface {
	Foo()
}

func (a *aRepository) Foo() {
	panic("needs implementation")
}

func InitARepository(a model.A) ARepository {
	aRepo := new(aRepository)
	aRepo.repo = a
	return aRepo
}
