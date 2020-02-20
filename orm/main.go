package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// User orm model
type User struct {
	gorm.Model
	Name  string
	Email string
}

func initialMigration() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect db")
	}
	defer db.Close()

	// migrate schema
	db.AutoMigrate(&User{})
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "all users endpoint hit")

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect db")
	}
	defer db.Close()

	var users []User
	db.Find(&users)

	fmt.Println("{}", users)

	json.NewEncoder(w).Encode(users)
}

func newUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "new user endpoint hit")

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect db")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	user := User{Name: name, Email: email}
	db.Create(&user)
	fmt.Fprintf(w, "user created")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "delete user endpoint hit")

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect db")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	//email := vars["email"]

	var user User
	db.Where("name = ?", name).Find(&user)

	db.Delete(&user)

	fmt.Fprintf(w, "user deleted")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "update user endpoint hit")

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect db")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	var user User
	db.Where("name = ?", name).Find(&user)

	user.Email = email

	db.Save(&user)
	fmt.Fprintf(w, "user updated")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/users", allUsers).Methods("GET")
	myRouter.HandleFunc("/user/{name}", deleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}", updateUser).Methods("PUT")
	myRouter.HandleFunc("/user/{name}/{email}", newUser).Methods("POST")

	err := http.ListenAndServe(":9000", myRouter)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	fmt.Println("go ORM tut")
	initialMigration()
	handleRequests()
}
