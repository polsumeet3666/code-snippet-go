// go web examples - json

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// User struct
type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

func main() {

	http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
		peter := User{
			Firstname: "Elon",
			Lastname:  "Musk",
			Age:       48,
		}

		json.NewEncoder(w).Encode(peter)
	})

	http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
		var user User
		json.NewDecoder(r.Body).Decode(&user)

		fmt.Fprintf(w, " %s %s is %d year old", user.Firstname, user.Lastname, user.Age)
	})

	http.ListenAndServe(":8000", nil)

}

// $ go run json.go

// $ curl -s -XPOST -d'{"firstname":"Elon","lastname":"Musk","age":48}' http://localhost:8080/decode
// Elon Musk is 48 years old!

// $ curl -s http://localhost:8080/encode
// {"firstname":"John","lastname":"Doe","age":25}
