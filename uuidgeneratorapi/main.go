package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/google/uuid"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/uuid", handleUUIDGen).Methods("GET")
	log.Print("uuid generator api up and running on 9000")
	http.ListenAndServe(":9000", router)
}

func handleUUIDGen(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	str := generateUUID()
	ctx = context.WithValue(ctx, "txnId", str)
	r.WithContext(ctx)
	fmt.Fprint(w)
}

func generateUUID() string {

	uuidWithHyphen := uuid.New()
	ss := string(uuidWithHyphen)
	return fmt.Sprint(uuidWithHyphen)
}

func generateUUIDWithoutHyphen() string {
	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	return uuid
}
