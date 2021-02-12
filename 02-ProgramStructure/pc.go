package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"gopl/02-ProgramStructure/popcount"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	x, _ := strconv.Atoi(input.Text())
	fmt.Println(popcount.PopCount(uint64(x)))
}
