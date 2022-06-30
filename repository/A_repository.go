package repository

import "github.com/bagasfathoni/go-clean-architecture-template/model/entity"

type aRepository struct {
	repo entity.A
}

type ARepository interface {
	Foo()
}

func (a *aRepository) Foo() {
	panic("needs implementation")
}

func InitARepository(a entity.A) ARepository {
	aRepo := new(aRepository)
	aRepo.repo = a
	return aRepo
}
