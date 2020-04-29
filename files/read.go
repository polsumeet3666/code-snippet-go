package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	byte, err := ioutil.ReadFile("sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	const Header = "MOBILE_NO|FREE_TEXT"
	//	fmt.Println(string(byte))

	byteReader := bytes.NewReader(byte)
	bufReader := bufio.NewReader(byteReader)
	count := 0
	for line, _, _ := bufReader.ReadLine(); line != nil; {
		//fmt.Println(string(line))
		// if Header == string(line) {
		// 	fmt.Println("same")
		// }
		//break
		line, _, _ = bufReader.ReadLine()
		result := strings.Split(string(line), "|")
		fmt.Println(result[0][0:10])
		stringMob := result[0][0:10]
		ii, err := strconv.Atoi(stringMob)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(ii)
		//stringSep := result[0][10:11]
		//fmt.Println(result[0][10:11])

		//strings.
		fmt.Println(result)
		break

	}
	fmt.Println(count)

}

func checkContent(line []byte) error {

	lineStr := string(line)
	recordSplit := strings.Split(lineStr, "|")
	if len(recordSplit) != 2 {
		return errors.New("Invalid content delimiter")
	}

	mobileStr := recordSplit[0][0:10]
	_, err := strconv.Atoi(mobileStr)
	if err != nil {
		return errors.New("Invalid mobile number")
	}

	return nil

}
