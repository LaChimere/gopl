package main

import (
    "fmt"
    "os"
)

func main() {
    for idx, arg := range os.Args[:] {
        // fmt.Printf("%d: %s\n", idx, arg)
        fmt.Println(idx, arg)
    }
}

