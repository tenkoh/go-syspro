package main

import (
	"encoding/csv"
	"os"
)

var toWrite = [][]string{
	{"name", "age"},
	{"john", "18"},
}

func main() {
	file, err := os.Create("tmp.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	for _, t := range toWrite {
		writer.Write(t)
	}
	writer.Flush()
}
