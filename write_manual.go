package main

import (
	"io"
	"log"
	"os"
)

func main4() {
	row, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer row.Close()

	write, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer write.Close()

	n, err := io.Copy(write, row)
	if err != nil {
		panic(err)
	}

	log.Printf("Copied %v bytes\n", n)

	main2()
}
