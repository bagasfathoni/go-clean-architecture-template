package manager

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/bagasfathoni/go-clean-architecture-template/config"
	_ "github.com/lib/pq"
)

type infra struct {
	db     *sql.DB
	config config.Config
}

type InfraManager interface {
	SqlDb() *sql.DB
}

// SqlDb implements InfraManager
func (i *infra) SqlDb() *sql.DB {
	return i.db
}

func initDbResource(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	} else {
		fmt.Println("Connection Success!")
	}
	return db, nil
}

func InitInfra(config *config.Config) InfraManager {
	dbResource, err := initDbResource(config.DataSoruceName)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &infra{db: dbResource}
}
