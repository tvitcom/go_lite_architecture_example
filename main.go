package main

import (
	"database/sql"
	"log"
	"net/http"

	"my.localhost/others/bookstore/handlers/books"
	"my.localhost/others/bookstore/models"

	_ "github.com/go-sql-driver/mysql"
)

const (
	STORAGE_DRV = "mysql"
	STORAGE_DSN = "bookstore:pass_to_bookstore@/bookstore"
)

func main() {
	db, err := sql.Open(STORAGE_DRV, STORAGE_DSN)
	if err != nil {
		log.Fatal(err)
	}

	env := &models.SetupStorage{DB: db}

	http.Handle("/books", books.Index(env))
	http.ListenAndServe(":3000", nil)
}
