package main

import (
	"fmt"
	"log"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func main() {
	var conn *ssh.Client

	// client config
	var config = &ssh.ClientConfig{
		User: "scrubbing_sftp",
		Auth: []ssh.AuthMethod{
			ssh.Password("scrubbing_sftp"),
		},
	}
	fmt.Println("1")

	// connect
	conn, err := ssh.Dial("tcp", "10.160.217.51:22", config)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	fmt.Println("2")

	// open an SFTP session over an existing ssh connection.
	sftp, err := sftp.NewClient(conn)
	if err != nil {
		log.Fatal(err)
	}
	defer sftp.Close()
	fmt.Println("3")

	// walk a directory
	w := sftp.Walk("/home/user")
	for w.Step() {
		if w.Err() != nil {
			continue
		}
		log.Println(w.Path())
	}
	fmt.Println("4")

	// leave your mark
	f, err := sftp.Create("hello.txt")
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte("Hello world!")); err != nil {
		log.Fatal(err)
	}
	fmt.Println("5")

	// check it's there
	fi, err := sftp.Lstat("hello.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(fi)

	fmt.Println("red")
}
