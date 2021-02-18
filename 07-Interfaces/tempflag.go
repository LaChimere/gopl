package main

import (
	"flag"
	"fmt"
	"gopl/02-ProgramStructure/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
