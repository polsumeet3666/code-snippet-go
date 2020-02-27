// go web examples - hello world
package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello , you have requested %v", r.URL.Path)
	})

	http.ListenAndServe(":8000", nil)
}
