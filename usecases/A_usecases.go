package usecases

import (
	"fmt"
	"strings"

	"github.com/bagasfathoni/go-clean-architecture-template/model"
	"github.com/bagasfathoni/go-clean-architecture-template/repository"
)

type aUsecases struct {
	aRepo repository.ARepository
}

type AUsecases interface {
	CreateNewA(newA *model.A) error
	UpdateNameById(id int, newName string) error
	UpdateStatusById(id int) error
	DeleteById(id int) error
	GetById(id int) (model.A, error)
	GetAllWithTrueStatus() ([]model.A, error)
	GetByLookAlikeName(name string) ([]model.A, error)
}

func (a *aUsecases) CreateNewA(newA *model.A) error {
	err := a.aRepo.Create(newA)
	if err != nil {
		return fmt.Errorf("failed to create a new A with error: %s", err.Error())
	}
	return nil
}

func (a *aUsecases) DeleteById(id int) error {
	err := a.aRepo.DeleteBy(map[string]interface{}{"id": id})
	if err != nil {
		return fmt.Errorf("failed to delete with error: %s", err.Error())
	}
	return nil
}

func (a *aUsecases) GetAllWithTrueStatus() ([]model.A, error) {
	res, err := a.aRepo.FindAllBy(map[string]interface{}{"status": true}, "created_at desc")
	if err != nil {
		return nil, fmt.Errorf("failed to get result with error: %s", err.Error())
	}
	return res, nil
}

func (a *aUsecases) GetById(id int) (model.A, error) {
	res, err := a.aRepo.FindBy(map[string]interface{}{"id": id})
	if err != nil {
		return model.A{}, fmt.Errorf("failed to get result with error: %s", err.Error())
	}
	return res, nil

}

func (a *aUsecases) GetByLookAlikeName(name string) ([]model.A, error) {
	res, err := a.aRepo.FindAllUsingCustomQuery("LOWER(name) LIKE %%%s%%", strings.ToLower(name))
	if err != nil {
		return nil, fmt.Errorf("failed to get result with error: %s", err.Error())
	}
	return res, nil
}

func (a *aUsecases) UpdateNameById(id int, newName string) error {
	err := a.aRepo.UpdateBy(map[string]interface{}{"id": id}, map[string]interface{}{"name": newName})
	if err != nil {
		return fmt.Errorf("failed to update record with error: %s", err.Error())
	}
	return nil
}

func (a *aUsecases) UpdateStatusById(id int) error {
	res, err := a.aRepo.FindBy(map[string]interface{}{"id": id})
	if err != nil {
		return fmt.Errorf("failed to get result with error: %s", err.Error())
	}

	err = a.aRepo.UpdateBy(map[string]interface{}{"id": res.Id}, map[string]interface{}{"status": !res.Status})
	if err != nil {
		return fmt.Errorf("failed to update record with error: %s", err.Error())
	}
	return nil

}

func InitCustomerUsecases(a repository.ARepository) AUsecases {
	return &aUsecases{aRepo: a}
}
