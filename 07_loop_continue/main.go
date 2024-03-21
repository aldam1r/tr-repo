package main

import "fmt"

func main() {
	x := 0
	for {
		if x%2 == 0 {
			x++
			continue
		}
		if x > 20 {
			break
		}
		fmt.Println(x)
		x++
	}
}
