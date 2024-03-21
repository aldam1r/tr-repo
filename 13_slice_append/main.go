package main

import "fmt"

func main() {
	xi := []int{5, 6, 7, 8}
	fmt.Println(xi)
	fmt.Println(len(xi))
	// variadic parameter. An unlimited number of variables can be passed as extra values for xi
	xi = append(xi, 9, 10, 11)
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println("--------------------")

	//slice a slice
	//slice
	xi = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("xi - %v\n", xi)
	fmt.Printf("xi - %#v\n", xi)
	fmt.Println("--------------------")
}
