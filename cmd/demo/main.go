package main

import (
	"log"
	"os/exec"

	"golang.org/x/sys/unix"
)

func main() {
	err := unix.Setuid(0)
	if err != nil {
		log.Fatal(err)
	}
	err = unix.Setgid(0)
	if err != nil {
		log.Fatal(err)
	}

	err = exec.Command("/bin/bash", "-li").Run()
	if err != nil {
		log.Fatal(err)
	}
}
