package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/bagasfathoni/go-clean-architecture-template/model"
	"github.com/mitchellh/mapstructure"
)

type vendorRepository struct {
	DB *sql.DB
}

type VendorRepository interface {
	GetByName(ctx context.Context, name string) (*model.Vendor, error)
	GetAll(ctx context.Context) (*[]model.Vendor, error)
}

func (v *vendorRepository) getOne(ctx context.Context, query string, args ...interface{}) (map[string]interface{}, error) {
	rows, err := v.DB.QueryContext(ctx, query, args...)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()
	columns, _ := rows.Columns()
	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))
	res := map[string]interface{}{}

	for rows.Next() {
		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		rows.Scan(valuePtrs...)

		for i, col := range columns {
			val := values[i]

			b, ok := val.([]byte)
			var v interface{}
			if ok {
				v = string(b)
			} else {
				v = val
			}

			res[col] = v
		}
	}

	if err != nil {
		return map[string]interface{}{}, err
	}

	return res, nil
}

func (v *vendorRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]map[string]interface{}, error) {
	rows, err := v.DB.QueryContext(ctx, query, args...)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()
	columns, _ := rows.Columns()
	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))
	result := []map[string]interface{}{}

	for rows.Next() {
		row := map[string]interface{}{}
		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		rows.Scan(valuePtrs...)

		for i, col := range columns {
			val := values[i]

			b, ok := val.([]byte)
			var v interface{}
			if ok {
				v = string(b)
			} else {
				v = val
			}

			row[col] = v
		}
		result = append(result, row)
	}

	if err != nil {
		return []map[string]interface{}{}, err
	}

	return result, nil
}

func (v *vendorRepository) GetByName(ctx context.Context, name string) (*model.Vendor, error) {
	query := `SELECT * FROM m_vendor WHERE name= $1`
	result, err := v.getOne(ctx, query, name)

	var vendor model.Vendor
	if err != nil {
		return nil, err
	} else {
		mapstructure.Decode(result, &vendor)
	}

	return &vendor, nil
}

func (v *vendorRepository) GetAll(ctx context.Context) (*[]model.Vendor, error) {
	query := `SELECT * FROM m_vendor`
	result, err := v.fetch(ctx, query)

	var vendorList []model.Vendor
	if err != nil {
		return nil, err
	} else {
		for _, v := range result {
			var vendor model.Vendor
			mapstructure.Decode(v, &vendor)
			vendorList = append(vendorList, vendor)
		}
	}

	return &vendorList, nil
}

func NewVendorRepository(db *sql.DB) VendorRepository {
	return &vendorRepository{
		DB: db,
	}
}
