package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"slices"
	"testing"
)

type getBookmarkRes struct {
	Bookmark Bookmark `json:"bookmark"`
}

type getBookmarksRes struct {
	Bookmarks []Bookmark `json:"bookmarks"`
}

type getMessageResponse struct {
	Message string `json:"message"`
}

type notBookmark struct {
	NotTitle    string `json:"notTitle"`
	URL         string `json:"url"`
	Description string `json:"description"`
}

func TestGetBookmarks(t *testing.T) {
	handler := InitialBookmarkHandler()
	expBookmarks := handler.bookmarks

	t.Run("returns bookmarks", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/bookmarks", nil)
		w := httptest.NewRecorder()
		var bookmarksRes getBookmarksRes
		handler.getBookmarks(w, req)

		res := w.Result()
		err := json.NewDecoder(res.Body).Decode(&bookmarksRes)
		if err != nil {
			t.Errorf("error decoding json: %q", err.Error())
		}
		if res.StatusCode != http.StatusOK {
			t.Errorf("expected status 200, got %d", res.StatusCode)
		} else if !slices.EqualFunc(bookmarksRes.Bookmarks, expBookmarks, func(bmr, bm Bookmark) bool {
			return bmr == bm
		}) {
			t.Errorf("expected bookmarks %v, got %v", expBookmarks, bookmarksRes.Bookmarks)
		}
	})
	t.Run("returns no bookmarks message", func(t *testing.T) {
		handler = &BookmarkHandler{}
		req := httptest.NewRequest(http.MethodGet, "/bookmarks", nil)
		w := httptest.NewRecorder()
		var bookmarksRes getMessageResponse
		handler.getBookmarks(w, req)

		res := w.Result()
		err := json.NewDecoder(res.Body).Decode(&bookmarksRes)
		if err != nil {
			t.Errorf("error decoding json: %q", err.Error())
		}
		if res.StatusCode != http.StatusOK {
			t.Errorf("expected status 200, got %d", res.StatusCode)
		} else if bookmarksRes.Message != "no bookmarks saved yet :(" {
			t.Errorf("expected bookmarks %v, got %v", "no bookmarks saved yet :(", bookmarksRes.Message)
		}
	})
}

func TestGetBookmark(t *testing.T) {
	handler := InitialBookmarkHandler()
	router := http.NewServeMux()
	router.HandleFunc("/bookmarks/{id}", handler.getBookmark)
	var bookmarkRes getBookmarkRes
	var messageRes getMessageResponse

	t.Run("get a bookmark by id", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/bookmarks/1", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		res := w.Result()
		err := json.NewDecoder(res.Body).Decode(&bookmarkRes)
		if err != nil {
			err = json.NewDecoder(res.Body).Decode(&messageRes)
			if err != nil {
				t.Errorf("error decoding json: %q, response body :", err.Error())
			} else {
				t.Errorf("error message: %s", messageRes.Message)
			}
		}
		if res.StatusCode != http.StatusOK {
			t.Errorf("expected status 200 but got %d", res.StatusCode)
		} else if bookmarkRes.Bookmark != handler.bookmarks[0] {
			t.Errorf("expected bookmark %v but got %v", handler.bookmarks[0], bookmarkRes)
		}
	})

	t.Run("get a bookmark by id but dosen't exist", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/bookmarks/6", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		res := w.Result()
		err := json.NewDecoder(res.Body).Decode(&messageRes)
		if err != nil {
			err = json.NewDecoder(res.Body).Decode(&messageRes)
			if err != nil {
				t.Errorf("error decoding json: %q, response body :", err.Error())
			} else {
				t.Errorf("error message: %s", messageRes.Message)
			}
		}
		if res.StatusCode != http.StatusNotFound {
			t.Errorf("expected status 404 but got %d", res.StatusCode)
		} else if messageRes.Message != "no such bookmark id: 6" {
			t.Errorf(`expected error message: "%s" but got "%s"`, "no such bookmark id: 6", messageRes.Message)
		}
	})
}

func TestCreateBookmark(t *testing.T) {
	handler := InitialBookmarkHandler()

	t.Run("create a new bookmark", func(t *testing.T) {
		createdBookmark := BookmarkBase{
			Title:       "Les fleur du mal",
			URL:         "myroom.irl",
			Description: "Basically a french weirdo from the 19th century making horny poetry",
		}
		reqBody, err := json.Marshal(createdBookmark)
		if err != nil {
			log.Fatalf("impossible to marshal BookmarkBase: %s", err.Error())
		}
		req := httptest.NewRequest(http.MethodPost, "/bookmarks", bytes.NewReader(reqBody))
		w := httptest.NewRecorder()
		var messageRes getMessageResponse
		handler.createBookmark(w, req)
		res := w.Result()
		err = json.NewDecoder(res.Body).Decode(&messageRes)
		if err != nil {
			t.Errorf("error decoding json %s", err.Error())
		}
		if res.StatusCode != http.StatusOK {
			t.Errorf("expected status code 200 but got %d", res.StatusCode)
		}
	})
	t.Run("fail to create a bookmark", func(t *testing.T) {
		createdBookmark := notBookmark{
			NotTitle:    "Les fleur du mal",
			URL:         "myroom.irl",
			Description: "Basically a french weirdo from the 19th century making horny poetry",
		}
		reqBody, err := json.Marshal(createdBookmark)
		if err != nil {
			log.Fatalf("impossible to marshal BookmarkBase: %s", err.Error())
		}
		req := httptest.NewRequest(http.MethodPost, "/bookmarks", bytes.NewReader(reqBody))
		w := httptest.NewRecorder()
		var messageRes getMessageResponse
		handler.createBookmark(w, req)
		res := w.Result()
		err = json.NewDecoder(res.Body).Decode(&messageRes)
		if err != nil {
			t.Errorf("error decoding json %s", err.Error())
		}
		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status code 400 but got %d", res.StatusCode)
		}
	})
}
