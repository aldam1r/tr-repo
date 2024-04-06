package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func main() {
	// read book
	book, err := os.ReadFile("great-gatsby.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer book.Close()

	words := wordCount(book)

}

func wordCount(r io.Reader) map[string]int {
	// map for word count
	wc := make(map[string]int)

	// slice of strings for all the words
	xs := make([]string, 0, 1_000)

	// scanner. Initialize new var with a scanner type "bufio.NewScanner(r)"
	s := bufio.NewScanner(r)
	s.Split()

}
