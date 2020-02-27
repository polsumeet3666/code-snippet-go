// spa - single page application
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
)

type spaHandler struct {
	staticPath string
	indexPath  string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// if we failed to get absolute path , then send badrequest and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// prepend path with path to static directory
	path = filepath.Join(h.staticPath, path)

	// check if file exists at given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// file does not exists serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		// if we get error other than file not exists then we have internal server error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// other wise serve file requested by url
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

func main() {

	// router
	router := mux.NewRouter()

	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		// an example api hander
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	// spa config
	spa := spaHandler{staticPath: "build", indexPath: "index.html"}
	router.PathPrefix("/").Handler(spa)

	// server

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// good practice : enforce timeout for server you create
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
