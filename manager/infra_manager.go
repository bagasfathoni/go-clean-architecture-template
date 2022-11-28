package manager

import (
	"gorm.io/gorm"
)

type infra struct {
	db *gorm.DB
}

type InfraManager interface {
	SqlDb() *gorm.DB
}

func (i *infra) SqlDb() *gorm.DB {
	return i.db
}

func InitInfra(db *gorm.DB) InfraManager {
	return &infra{db: db}
}
