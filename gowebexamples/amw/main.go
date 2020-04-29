// go web examples - advanced middleware

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Middleware -
type Middleware func(http.HandlerFunc) http.HandlerFunc

// Logging all request with its path and time it took to process
func Logging() Middleware {

	// create new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {
		// define handlerfunc here
		return func(w http.ResponseWriter, r *http.Request) {
			// do middleware things
			log.Println("logging mw")
			start := time.Now()
			defer func() { log.Println(r.URL.Path, time.Since(start)) }()

			// call next middleware in chain
			f(w, r)
		}
	}
}

// Method ensure that url can only be requested with specific method,else return a 400 bad request
func Method(m string) Middleware {

	// create new middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// define handler here
		return func(w http.ResponseWriter, r *http.Request) {
			// do middleware things here
			log.Println("method mw")
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			// call next middleware
			f(w, r)
		}
	}
}

// Chain middleware
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

// Hello handler
func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello word")
}

func main() {

	http.HandleFunc("/", Chain(Hello, Method("GET"), Logging()))

	http.ListenAndServe(":8000", nil)
}
