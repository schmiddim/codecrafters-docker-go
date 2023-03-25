package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
)

// Usage: your_docker.sh run <image> <command> <arg1> <arg2> ...
func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	//fmt.Println("Logs from your program will appear here!")

	command := os.Args[3]
	args := os.Args[4:len(os.Args)]

	cmd := exec.Command(command, args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	writerStd := bufio.NewWriter(os.Stdout)
	writerError := bufio.NewWriter(os.Stderr)
	_, err = writerStd.ReadFrom(stdout)
	_, err = writerError.ReadFrom(stderr)
	if err != nil {
		return
	}
}
