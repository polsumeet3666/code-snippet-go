// go web examples - password hashing
package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword using bcrypt
func HashPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err

}

// CheckPassword using hash
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {

	password := "secret"

	hash, _ := HashPassword(password) // error ignored for sake of simplicity

	fmt.Println("password:", password)
	fmt.Println("hash :", hash)

	match := CheckPassword(password, hash)
	fmt.Println("match :", match)

}
