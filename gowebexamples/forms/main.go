// go web example - forms

package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// ContactDetails struct
type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func main() {

	// template loading using must anf parsefile methods
	tmpl := template.Must(template.ParseFiles("form.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		details := ContactDetails{
			Email:   r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}
		fmt.Println(details)
		tmpl.Execute(w, struct{ Success bool }{true})
	})

	http.ListenAndServe(":8000", nil)

}
