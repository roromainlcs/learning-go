package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	handler := &BookmarkHandler{}
	router.HandleFunc("POST /bookmarks", handler.createBookmark)
	router.HandleFunc("GET /bookmarks/{id}", handler.getBookmark)
	router.HandleFunc("PUT /bookmarks/{id}", handler.updateBookmark)
	router.HandleFunc("GET /bookmarks", handler.getBookmarks)
	router.HandleFunc("DELETE /bookmarks/{id}", handler.deleteBookmark)
	fmt.Printf("server starting at 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
