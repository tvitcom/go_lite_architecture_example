package main

import (
	"my.localhost/others/bookstore/internal/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/books", handler.BooksIndex)
	http.ListenAndServe(":3000", nil)
}
