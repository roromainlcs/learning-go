package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"strconv"
	"time"
)

type Bookmark struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	URL         string    `json:"url"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type BookmarkHandler struct {
	bookmarks []Bookmark
}

func (bms *BookmarkHandler) createBookmark(w http.ResponseWriter, r *http.Request) {
	var newBookmark Bookmark
	err := json.NewDecoder(r.Body).Decode(&newBookmark)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	bms.bookmarks = append(bms.bookmarks, newBookmark)
}

func (bms BookmarkHandler) getBookmark(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	idx := slices.IndexFunc(bms.bookmarks, func(bm Bookmark) bool { return bm.ID == id })
	if idx == -1 {
		http.Error(w, fmt.Sprintf("no such bookmark id %d", id), http.StatusBadRequest)
	}
	json.Marshal(bms.bookmarks[idx])
}

func (bms *BookmarkHandler) updateBookmark(w http.ResponseWriter, r *http.Request) {
	var updatedBookmark Bookmark
	err := json.NewDecoder(r.Body).Decode(&updatedBookmark)
	id, errID := strconv.Atoi(r.PathValue("id"))

	if errID != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	idx := slices.IndexFunc(bms.bookmarks, func(bm Bookmark) bool { return bm.ID == id })

	if idx == -1 {
		http.Error(w, fmt.Sprintf("no such bookmark id %q", id), http.StatusBadRequest)
	}

	bms.bookmarks[idx] = updatedBookmark
}

func (bms BookmarkHandler) getBookmarks(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("coucou\n")
	var response []byte
	var err error

	if len(bms.bookmarks) == 0 {
		fmt.Printf("no bookmarks\n")
		response, err = json.Marshal("no bookmarks saved yet :(")
	} else {
		response, err = json.Marshal(bms.bookmarks)
	}
	if err != nil {
		http.Error(w, "error loading json", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (bms *BookmarkHandler) deleteBookmark(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	idx := slices.IndexFunc(bms.bookmarks, func(bm Bookmark) bool { return bm.ID == id })
	if idx == -1 {
		http.Error(w, fmt.Sprintf("no such bookmark id %q", id), http.StatusBadRequest)
	}
	bms.bookmarks = slices.Delete(bms.bookmarks, idx, idx)
	json.Marshal(fmt.Sprintf("bookmark %q deleted", id))
}
