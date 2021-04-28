package main

import (
	"fmt"
	"log"
	"my.localhost/others/bookstore/internal/config"
	"my.localhost/others/bookstore/internal/models"
	"net/http"
)

const (
	STORAGE_DRV = "mysql"
	STORAGE_DSN = "bookstore:pass_to_bookstore@/bookstore"
)

func main() {
	db, err := models.NewDB(STORAGE_DRV, STORAGE_DSN)
	if err != nil {
		log.Panic(err)
	}
	service := &config.configService{db}

	fmt.Println("Start web server on :3000")
	http.HandleFunc("/", service.booksIndex)
	http.ListenAndServe(":3000", nil)
}

func (service *configService) booksIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	bks, err := service.db.AllBooks()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}
}
