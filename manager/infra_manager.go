package manager

import (
	"log"

	"github.com/bagasfathoni/go-clean-architecture-template/config"
	"github.com/bagasfathoni/go-clean-architecture-template/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type infra struct {
	db     *gorm.DB
	config config.Config
}

type InfraManager interface {
	SqlDb() *gorm.DB
}

func (i *infra) SqlDb() *gorm.DB {
	return i.db
}

func InitInfra(config *config.Config) InfraManager {
	resource, err := initDbResource(config.DataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &infra{db: resource}
}

func initDbResource(dataSourceName string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	db.AutoMigrate(
		&model.A{},
	)
	if err != nil {
		return nil, err
	}
	return db, nil
}
