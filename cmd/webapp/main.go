package main

import (
	"fmt"
	"my.localhost/others/bookstore/internal/handler"
	"net/http"
)

func main() {
	fmt.Println("Default route: GET /books")
	http.HandleFunc("/books", handler.BooksIndex)
	http.ListenAndServe(":3000", nil)
}
