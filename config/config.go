package config

import (
	"fmt"
	"os"
)

type Config struct {
	DbConfig
	ApiConfig
}

type DbConfig struct {
	DataSoruceName string
}

type ApiConfig struct {
	Url string
}

func InitConfig() Config {
	return Config{DbConfig: initDbConfig(), ApiConfig: initApiConfig()}
}

func initDbConfig() DbConfig {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPass, dbName, dbPort)
	return DbConfig{DataSoruceName: dsn}
}

func initApiConfig() ApiConfig {
	url := os.Getenv("API_URL")
	return ApiConfig{Url: url}
}
