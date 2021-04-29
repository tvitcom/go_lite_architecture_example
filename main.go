package main

import (
	"my.localhost/others/bookstore/handler"
	"net/http"
)

/*
	"bookstore.alexedwards.net/config"
	"bookstore.alexedwards.net/models/books"
*/

func main() {
	http.HandleFunc("/books", handler.BooksIndex)
	http.ListenAndServe(":3000", nil)
}
