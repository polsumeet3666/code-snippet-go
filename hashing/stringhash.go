package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	str := "sample text for hashingsample text for hashingsample text for hashingsample text for hashingsample text for hashing"

	x := md5.Sum([]byte(str))
	aa := fmt.Sprintf("%x",x)
	fmt.Printf(aa)
}
