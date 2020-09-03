package main

import (
	"log"
	"os"
)

var cwd string

func main() {
	cwd, err := os.Getwd() // NOTE: wrong!
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	log.Printf("Working directory = %s", cwd)
}
