package handler

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"my.localhost/others/bookstore/config"
	"my.localhost/others/bookstore/models/books"
	"net/http"
)

func BooksIndex(w http.ResponseWriter, r *http.Request) {
	bks, err := books.All(config.DB)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}
}
