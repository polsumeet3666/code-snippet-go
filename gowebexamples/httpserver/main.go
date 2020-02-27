// go web examples - http server
package main

import (
	"fmt"
	"net/http"
)

func main() {
	// simple index route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to http server website")
	})

	// serve static files
	// create file server
	fs := http.FileServer(http.Dir("static/"))

	// stripPrefix map url to alternate dir on above fs
	http.Handle("/staticf/", http.StripPrefix("/staticf/", fs))

	http.ListenAndServe(":8000", nil)

}
