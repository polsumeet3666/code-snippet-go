package main

import (
	"encoding/json"
	"net/http"

	_ "./docs"
)

func main() {

	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	http.ListenAndServe(":8000", nil)

}
