package db

import (
	"log"
	"warunggpt-core-service/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func InitDB(cfg *config.Config) {
	var err error
	db, err = sqlx.Connect("mysql", cfg.DB.DataSourceName())
	if err != nil {
		log.Fatal(err)
	}
}

func GetDB() *sqlx.DB {
	return db
}
