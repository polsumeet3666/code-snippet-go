package main

import (
	"fmt"
	"strings"
)

func main() {
	ext := "application/vnd.ms-excel"
	allowed := "text/csv,application/vnd.ms-excel,application/zip"

	if index := strings.Index(allowed, ext); index == -1 {
		fmt.Println("not allowed")
	}
}
