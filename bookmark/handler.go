package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"strconv"
	"time"
)

type BookmarkBase struct {
	Title       string `json:"title"`
	URL         string `json:"url"`
	Description string `json:"description"`
}

type Bookmark struct {
	ID int `json:"id"`
	BookmarkBase
	CreatedAt time.Time `json:"created_at"`
}

type BookmarkHandler struct {
	bookmarks []Bookmark
}

func (bms *BookmarkHandler) createBookmark(w http.ResponseWriter, r *http.Request) {
	var receivedBookmark BookmarkBase
	err := json.NewDecoder(r.Body).Decode(&receivedBookmark)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newBookmarkID := bms.bookmarks[len(bms.bookmarks)-1].ID + 1
	newBookmark := Bookmark{BookmarkBase: receivedBookmark, ID: newBookmarkID, CreatedAt: time.Now()}
	bms.bookmarks = append(bms.bookmarks, newBookmark)
	Send(w, "bookmark created succesfully")
}

func (bms *BookmarkHandler) getBookmark(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	idx := slices.IndexFunc(bms.bookmarks, func(bm Bookmark) bool { return bm.ID == id })
	if idx == -1 {
		http.Error(w, fmt.Sprintf("no such bookmark id %d", id), http.StatusBadRequest)
		return
	}
	Send(w, bms.bookmarks[idx])
}

func (bms *BookmarkHandler) updateBookmark(w http.ResponseWriter, r *http.Request) {
	var updatedBookmark BookmarkBase
	err := json.NewDecoder(r.Body).Decode(&updatedBookmark)
	id, errID := strconv.Atoi(r.PathValue("id"))

	if errID != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	idx := slices.IndexFunc(bms.bookmarks, func(bm Bookmark) bool { return bm.ID == id })

	if idx == -1 {
		http.Error(w, fmt.Sprintf("no such bookmark id %d", id), http.StatusBadRequest)
		return
	}

	bms.bookmarks[idx].Title = updatedBookmark.Title
	bms.bookmarks[idx].Description = updatedBookmark.Title
	bms.bookmarks[idx].Title = updatedBookmark.Description
	bms.bookmarks[idx].CreatedAt = time.Now()
	Send(w, fmt.Sprintf("Bookmark of ID %d updated successfuly", id))
}

func (bms *BookmarkHandler) getBookmarks(w http.ResponseWriter, r *http.Request) {
	if len(bms.bookmarks) == 0 {
		fmt.Printf("no bookmarks\n")
		Send(w, "no bookmarks saved yet :(")
	}
	Send(w, bms.bookmarks, "bookmarks")
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
