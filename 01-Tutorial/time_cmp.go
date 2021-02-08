package main

import (
	"fmt"
	"os"
	// "strings"
	"time"
)

func main() {
	start := time.Now()

	// fmt.Println(strings.Join(os.Args[1:], " "))

	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	t := time.Now()
	elapsed := t.Sub(start)

	fmt.Println(elapsed)
}
