// go web example- logging middleware lmw
package main

import (
	"fmt"
	"log"
	"net/http"
)

// logging middleware
func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

// handler for /foo
func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "/foo")
}

// handler for /bar
func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "/bar")
}

func main() {

	// wrapping foo with logging middleware
	http.HandleFunc("/foo", logging(foo))

	// wrapping foo with logging middleware
	http.HandleFunc("/bar", logging(bar))

	// server
	http.ListenAndServe(":8000", nil)
}
