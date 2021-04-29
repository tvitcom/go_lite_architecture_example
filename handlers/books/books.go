package books

import (
	"fmt"
	"log"
	"net/http"

	"my.localhost/others/bookstore/models"
	"my.localhost/others/bookstore/models/book"
)

func Index(env *models.SetupStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bks, err := book.AllBooks(env.DB)
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		for _, bk := range bks {
			fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
		}
	}
}
