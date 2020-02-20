package main

import (
	"fmt"
	"os/exec"
	"runtime"

	log "github.com/sirupsen/logrus"
)

func execute() {

	// simple command without arguments
	out, err := exec.Command("ls").Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))

	// simple command without arguments
	out, err = exec.Command("pwd").Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))

	// simple command with arguments
	out, err = exec.Command("ls", "-ltr").Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))
}

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("cannot run these commands on windows")
	} else {
		fmt.Println(runtime.GOOS)
		execute()
	}
}
