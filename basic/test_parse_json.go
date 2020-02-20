package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	jsonFile, err := os.Open("users.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("file open successfully")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	fmt.Println(result["users"])
}
