package main

import (
	"io/ioutil"
	"log"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func main() {

	startTime := time.Now()
	addr := "10.160.217.51:22"
	config := &ssh.ClientConfig{
		User: "scrubbing_sftp",
		Auth: []ssh.AuthMethod{
			ssh.Password("scrubbing_sftp"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		panic("Failed to dial: " + err.Error())
	}
	client, err := sftp.NewClient(conn)
	if err != nil {
		panic("Failed to create client: " + err.Error())
	}
	// Close connection
	defer client.Close()
	cwd, err := client.Getwd()
	println("Current working directory:", cwd)

	srcPath := "C:\\Users\\tejendra.singhadiya\\Desktop\\testprogram\\test2java.txt"
	dstPath := "/input/"
	filename := "TestFile0.txt"

	CreateFile, err := client.Create(cwd + dstPath + filename)
	if err != nil {
		log.Fatal(err)
	}
	println(CreateFile)

	bytes, err := ioutil.ReadFile(srcPath)
	if err != nil {
		log.Fatal(err)
	}

	fileSize, err := CreateFile.Write(bytes)
	if err != nil {
		log.Fatal(err)
	}
	println(fileSize)

	println(time.Since(startTime) / 1000000)

}
