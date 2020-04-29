// go web exmaples - gorilla/mux

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// create new mux router
	router := mux.NewRouter()

	// custom handlerFunc to handle segments/ dynamic parameters
	router.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		
		
		
		// get url variables
		vars := mux.Vars(r)
		fmt.Println(vars)
		
		title := vars["title"]
		page := vars["page"]
		fmt.Println(title,page)
		fmt.Fprintf(w, " you have request book : %v and page : %v", title, page)
	})

	// set mux router to server
	http.ListenAndServe(":8000", router)

}
