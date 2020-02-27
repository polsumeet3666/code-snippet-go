package main

import (
	"html/template"
	"log"
	"net/http"
)

// Todo struct
type Todo struct {
	Title string
	Done  bool
}

// TodoPageData struct
type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	tmpl, err := template.ParseFiles("layout.html")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My Sample Todos",
			Todos: []Todo{
				{Title: "task 1 ", Done: false},
				{Title: "task 2 ", Done: false},
				{Title: "task 3 ", Done: true},
			},
		}

		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":8000", nil)

}
