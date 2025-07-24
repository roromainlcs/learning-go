package main

import (
	"encoding/json"
	"net/http"
)

func Send(w http.ResponseWriter, response any, name ...string) {
	key := "message"

	if len(name) > 0 && name[0] != "" {
		key = name[0]
	}
	sent, err := json.Marshal(map[string]any{key: response})

	if err != nil {
		http.Error(w, "error loading json", http.StatusInternalServerError)
	}
	w.Write(sent)
}
