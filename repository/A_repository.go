package repository

import (
	"github.com/bagasfathoni/go-clean-architecture-template/model"
	"gorm.io/gorm"
)

/*
	Repository store basic methods that will directly hit to the database.
	In this github branch we create basic CRUD methods using GORM
*/

type aRepository struct {
	db *gorm.DB
}

type ARepository interface {
	Create(newA *model.A) error
	UpdateBy(by map[string]interface{}, value map[string]interface{}) error
	FindBy(by map[string]interface{}) (model.A, error)
	FindAllBy(by map[string]interface{}, orderBy string) ([]model.A, error)
	FindAllByWithPagination(by map[string]interface{}, page, itemPerPage int, orderBy string) ([]model.A, error)
	FindAllUsingCustomQuery(query, orderBy string) ([]model.A, error)
	FindAllUsingCustomQueryWithPagination(query, orderBy string, page, itemPerPage int) ([]model.A, error)
	Delete(A *model.A) error
}

func (a *aRepository) Create(transferOrder *model.A) error {
	result := a.db.Create(&transferOrder).Error
	return result
}

func (a *aRepository) FindBy(by map[string]interface{}) (model.A, error) {
	var res model.A
	err := a.db.Where(by).First(&res).Error
	return res, err

}

func (a *aRepository) FindAllByWithPagination(by map[string]interface{}, page, itemPerPage int, orderBy string) ([]model.A, error) {
	var res []model.A
	err := a.db.Where(by).Order(orderBy).Limit(itemPerPage).Offset(page).Find(&res).Error
	return res, err

}

func (a *aRepository) FindAllBy(by map[string]interface{}, orderBy string) ([]model.A, error) {
	var res []model.A
	err := a.db.Where(by).Order(orderBy).Unscoped().Find(&res).Error
	return res, err

}

func (a *aRepository) FindAllUsingCustomQuery(query, orderBy string) ([]model.A, error) {
	var res []model.A
	err := a.db.Where(query).Order(orderBy).Unscoped().Find(&res).Error
	return res, err
}

func (a *aRepository) FindAllUsingCustomQueryWithPagination(query, orderBy string, page, itemPerPage int) ([]model.A, error) {
	var res []model.A
	err := a.db.Where(query).Order(orderBy).Limit(itemPerPage).Offset(page).Find(&res).Error
	return res, err
}

func (a *aRepository) UpdateBy(by map[string]interface{}, value map[string]interface{}) error {
	return a.db.Model(model.A{}).Where(by).Updates(value).Error
}

func (a *aRepository) Delete(A *model.A) error {
	res := a.db.Delete(A).Error
	return res
}

func InitARepository(a *gorm.DB) ARepository {
	return &aRepository{db: a}
}
