package config

import (
	"fmt"
	"os"
)

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

func LoadConfig() *Config {
	return &Config{
		DB: &DBConfig{
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Name:     os.Getenv("DB_NAME"),
		},
	}
}

func (db *DBConfig) DataSourceName() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		db.User, db.Password, db.Host, db.Port, db.Name)
}
