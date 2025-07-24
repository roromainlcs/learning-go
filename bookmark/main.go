package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	router := http.NewServeMux()
	handler := InitialBookmarkHandler()
	router.HandleFunc("POST /bookmarks", handler.createBookmark)
	router.HandleFunc("GET /bookmarks/{id}", handler.getBookmark)
	router.HandleFunc("PUT /bookmarks/{id}", handler.updateBookmark)
	router.HandleFunc("GET /bookmarks", handler.getBookmarks)
	router.HandleFunc("DELETE /bookmarks/{id}", handler.deleteBookmark)
	fmt.Printf("server starting at 8080")
	log.Fatal(http.ListenAndServe(":8080", Middleware(router)))
}

func InitialBookmarkHandler() *BookmarkHandler {
	return &BookmarkHandler{
		bookmarks: []Bookmark{
			{
				ID: 1,
				BookmarkBase: BookmarkBase{
					Title:       "Resister a la culpabilisation",
					Description: "Feminist book from mona cholet about the culpabilisation of the minorities",
					URL:         "myRoom.irl",
				},
				CreatedAt: time.Now(),
			},
			{
				ID: 2,
				BookmarkBase: BookmarkBase{
					Title:       "Berserk",
					Description: "RIP Kentaro Miura",
					URL:         "myRoom.irl",
				},
				CreatedAt: time.Now(),
			},
			{
				ID: 3,
				BookmarkBase: BookmarkBase{
					Title:       "SCUM",
					Description: "A milestone of feminism literature by Valerie Solanas",
					URL:         "myRoom.irl",
				},
				CreatedAt: time.Now(),
			},
			{
				ID: 4,
				BookmarkBase: BookmarkBase{
					Title:       "L'histoire de ta betise",
					Description: "Political book about 2017 french presidential election by Francois Begaudeau",
					URL:         "myRoom.irl",
				},
				CreatedAt: time.Now(),
			},
		},
	}
}
