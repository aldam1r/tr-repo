package main

import "fmt"

func main() {
	i := 0
	for i < 20 {
		fmt.Printf("Value of i is of type %T and has a value of %v\n", i, i)
		i++
	}
}
