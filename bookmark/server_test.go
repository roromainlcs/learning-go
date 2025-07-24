package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"slices"
	"testing"
)

type getBookmarksRes struct {
	Bookmarks []Bookmark `json:"bookmarks"`
}

type getEmptyBookmarksRes struct {
	Message string `json:"message"`
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
		var bookmarksRes getEmptyBookmarksRes
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
