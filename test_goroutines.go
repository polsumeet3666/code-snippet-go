package main

import (
	"fmt"
	"time"
)

func compute(value int) {
	for i := 0; i < value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}

func main() {
	fmt.Println("go routine example")

	go compute(10)
	go compute(10)

	var input string
	fmt.Scanln(&input)
}
