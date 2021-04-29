package config

import (
	"database/sql"
	"log"
)

const (
	STORAGE_DRV = "mysql"
	STORAGE_DSN = "bookstore:pass_to_bookstore@/bookstore"
)

var DB *sql.DB

func init() {
	var err error

	DB, err = sql.Open(STORAGE_DRV, STORAGE_DSN)
	if err != nil {
		log.Fatal(err)
	}
}
