// go web examples - sessions
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func secret(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// check for auth
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}
	fmt.Fprintln(w, "cake is lie")
}

func login(w http.ResponseWriter, r *http.Request) {

	// get session store
	sessions, _ := store.Get(r, "cookie-name")

	// authenticate her

	// set user authenticated
	sessions.Values["authenticated"] = true
	sessions.Save(r, w)
}

func logout(w http.ResponseWriter, r *http.Request) {

	// get store
	session, _ := store.Get(r, "cookie-name")

	// revoke here

	session.Values["authenticated"] = false
	session.Save(r, w)

}

func main() {

	http.HandleFunc("/login", login)
	http.HandleFunc("/secret", secret)
	http.HandleFunc("/logout", logout)

	http.ListenAndServe(":8000", nil)
}

// $ go run sessions.go

// $ curl -s http://localhost:8080/secret
// Forbidden

// $ curl -s -I http://localhost:8080/login
// Set-Cookie: cookie-name=MTQ4NzE5Mz...

// $ curl -s --cookie "cookie-name=MTQ4NzE5Mz..." http://localhost:8080/secret
// The cake is a lie
