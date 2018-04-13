package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, error := os.Open(os.Args[1])

	if error != nil {
		fmt.Println("error occurred trying to read file", error)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
