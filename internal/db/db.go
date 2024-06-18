package db

import (
	"database/sql"
	"log"
	"warunggpt-core-service/config"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB(cfg *config.Config) {
	dsn := cfg.DB.DataSourceName()

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	log.Println("Database connected!")

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
}

func GetDB() *sql.DB {
	return db
}
