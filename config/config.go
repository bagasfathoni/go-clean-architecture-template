package config

import (
	"fmt"
	"os"
)

type Config struct {
	ApiConfig
	DbConfig
}

func InitConfig() Config {
	return Config{ApiConfig: initApiConfig(), DbConfig: initDbConfig()}
}

// DB Config
type DbConfig struct {
	DataSourceName string
}

func initDbConfig() DbConfig {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPass, dbName, dbPort)
	return DbConfig{DataSourceName: dsn}
}

// REST API Config
type ApiConfig struct {
	Url string
}

func initApiConfig() ApiConfig {
	api := os.Getenv("API_URL")
	return ApiConfig{Url: api}

}
