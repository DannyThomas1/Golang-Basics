package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := os.Args[1]

	readFile(fileName)
}

func readFile(filename string) {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Failed to read file")
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
