package main

import (
	"log"
	"os"
)

var cwd string

// func init() {
// 	cwd, err := os.Getwd() // compile error: unused: cwd
// 	if err != nil {
// 		log.Fatalf("os.Getwd failed: %v", err)
// 	}
// }

// func init() {
// 	cwd, err := os.Getwd() // NOTE: wrong!
// 	if err != nil {
// 		log.Fatalf("os.Getwd failed: %v", err)
// 	}
// 	log.Printf("Working directory = %s", cwd)
// }

func init() {
	var err error
	cwd, err = os.Getwd() // NOTE: right!
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
}

func main() {
	log.Printf("Working directory = %s", cwd)
}
