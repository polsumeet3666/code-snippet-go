package main

import (
	"fmt"
	"os"
)

func main() {

	f, _ := os.Create("sample2.txt")
	defer f.Close()

	for i := 0; i < 500000; i++ {
		f.WriteString(string(i) + "\n")
//		f.Sync()
	}
	fmt.Println("done")
}
