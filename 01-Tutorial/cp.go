package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "usage: %s src dest\n", os.Args[0])
		os.Exit(1)
	}

	src, dest := os.Args[1], os.Args[2]
	data, err := ioutil.ReadFile(src)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ReadFile: %v\n", err)
		os.Exit(1)
	}

	err = ioutil.WriteFile(dest, data, 0664)
	if err != nil {
		fmt.Fprintf(os.Stderr, "WriteFile: %v\n", err)
		os.Exit(1)
	}
}
