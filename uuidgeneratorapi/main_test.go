package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
)

func TestApi(t *testing.T) {
	response, err := http.Get("http://localhost:9000/api/v1/generateuuid")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// print data on console
	fmt.Println(string(responseData))

}
