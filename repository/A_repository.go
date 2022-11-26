/*
Repository store basic methods that will directly hit to the database.
In this github branch we create basic CRUD methods using GORM
*/
package repository

import (
	"github.com/bagasfathoni/go-clean-architecture-template/model"
	"gorm.io/gorm"
)

type aRepository struct {
	db *gorm.DB
}

type ARepository interface {
	Create(newA *model.A) error
	UpdateBy(by map[string]interface{}, value map[string]interface{}) error
	DeleteBy(map[string]interface{}) error
	FindBy(by map[string]interface{}) (model.A, error)
	FindAllBy(by map[string]interface{}, orderBy string) ([]model.A, error)
	FindAllByWithPagination(by map[string]interface{}, page, itemPerPage int, orderBy string) ([]model.A, error)
	FindAllUsingCustomQuery(query, orderBy string) ([]model.A, error)
	FindAllUsingCustomQueryWithPagination(query, orderBy string, page, itemPerPage int) ([]model.A, error)
}

/* Create/insert a new record to the table by defining the model explicitly */
func (a *aRepository) Create(transferOrder *model.A) error {
	result := a.db.Create(&transferOrder).Error
	return result
}

/*
Find a record by using the row name and the data.

	Example:
	FindBy(map[string]interface{}{"id":1}) // will get result with ID = 1
*/
func (a *aRepository) FindBy(by map[string]interface{}) (model.A, error) {
	var res model.A
	err := a.db.Where(by).First(&res).Error
	return res, err

}

/*
Find any records by using the row name and the data.

	Example:
	FindAllBy(map[string]interface{}{}, "created_at desc") // will get all result in the table
	FindAllBy(map[string]interface{}{"name":"XXX"}, "created_at desc") // will get all result with name = "XXX"
*/
func (a *aRepository) FindAllBy(by map[string]interface{}, orderBy string) ([]model.A, error) {
	var res []model.A
	err := a.db.Where(by).Order(orderBy).Unscoped().Find(&res).Error
	return res, err

}

/*
Find any records by using the row name and the data. This will limit the result to specific number.

	Example:
	FindAllByWithPagination(map[string]interface{}{}, 1, 10, "created_at desc") // will get the first 10 result in the column
	FindAllByWithPagination(map[string]interface{}{}, 2, 10, "created_at desc") // will get the next 10 result in the column
	FindAllByWithPagination(map[string]interface{}{"name":"XXX"}, 2, 10, "created_at desc") // will get the first 10 result in the column with name = "XXX"
*/
func (a *aRepository) FindAllByWithPagination(by map[string]interface{}, page, itemPerPage int, orderBy string) ([]model.A, error) {
	var res []model.A
	err := a.db.Where(by).Order(orderBy).Limit(itemPerPage).Offset(page).Find(&res).Error
	return res, err

}

/*
Find any records by using custom SQL Query.

	Example:
	FindAllUsingCustomQuery("name = 'XXX' AND email == 'YYY'", "created_at desc") // will get all result in the column with name = "XXX" and email = "YYY"
*/
func (a *aRepository) FindAllUsingCustomQuery(query, orderBy string) ([]model.A, error) {
	var res []model.A
	err := a.db.Where(query).Order(orderBy).Unscoped().Find(&res).Error
	return res, err
}

/*
Find any records by using any SQL Query. This will limit the result to specific number.

	Example:
	FindAllUsingCustomQuery("name = 'XXX' AND email = 'YYY'", "created_at desc", 1, 10) // will get the first 10 result in the column with name = "XXX" and email = "YYY"
	FindAllUsingCustomQuery("name = 'XXX' AND email = 'YYY'", "created_at desc", 2, 10) // will get the next 10 result in the column with name = "XXX" and email = "YYY"
*/
func (a *aRepository) FindAllUsingCustomQueryWithPagination(query, orderBy string, page, itemPerPage int) ([]model.A, error) {
	var res []model.A
	err := a.db.Where(query).Order(orderBy).Limit(itemPerPage).Offset(page).Find(&res).Error
	return res, err
}

/*
Update a value in a record.

	Example:
	UpdateBy(map[string]interface{}{"id":1}, map[string]interface{}{"name":"YYY"}) // will update a record name with ID = 1 to "YYY"
*/
func (a *aRepository) UpdateBy(by map[string]interface{}, value map[string]interface{}) error {
	return a.db.Model(model.A{}).Where(by).Updates(value).Error
}

/*
Delete a record.

	Example:
	DeleteBy(map[string]interface{}{"id":1}) // will delete a record name with ID = 1
*/
func (a *aRepository) DeleteBy(by map[string]interface{}) error {
	return a.db.Model(model.A{}).Delete(by).Error
}

func InitARepository(a *gorm.DB) ARepository {
	return &aRepository{db: a}
}
