package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.CreateTemp("", "")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Fprintf(file, "%s %d %f", "hello", 1, 1.00)
}
