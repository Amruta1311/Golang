package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// fmt.Println(os.Args[1])	// Accesses the file
	f, err := os.Open(os.Args[1]) // returns a pointer to the file and error if occurs

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
