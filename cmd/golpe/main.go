package main

import (
	"log"
	"time"

	golpe "github.com/jm33-m0/go-lpe"
)

func main() {
	golpe.RunAll()
	for {
		time.Sleep(1 * time.Second)
		log.Println("sleeping...")
	}
}
