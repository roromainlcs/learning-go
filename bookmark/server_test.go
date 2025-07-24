package main

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// )

// func TestGetBookmarks(t *testing.T) {

// 	t.Run("get bookmarks", func(t *testing.T) {
// 		request, error := http.NewRequest(http.MethodGet, "/bookmarks", nil)
// 		response := httptest.NewRecorder()

// 		if error != nil {
// 			t.Errorf("error accessing server: %s", error.Error())
// 		}

// 		assertResponseBody(t, response.Body.String(), "want")
// 	})
// }

// func assertResponseBody(t testing.TB, got, want string) {
// 	t.Helper()

// 	if got != want {
// 		t.Errorf("got %q, want %q", got, want)
// 	}
// }
