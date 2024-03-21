package main

import "fmt"

func main() {
	x := 0
	for {
		if x > 20 {
			fmt.Println("Break!!!")
			break
		}
		fmt.Println("No break yet")
		x++
	}
}
