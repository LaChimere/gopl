package main

import (
	"log"
	"time"
)

func bigSlowOperation() {
	defer trace("1")()
	defer trace("2")()
	defer trace("3")()
	// ...lots of work...
	time.Sleep(1 * time.Second) // simulate slow operation
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func main() {
	bigSlowOperation()
}
