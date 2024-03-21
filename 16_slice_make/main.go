package main

import "fmt"

func main() {
	//Make a slice
	//slice
	xi := make([]int, 0, 10)

	fmt.Printf("xi - %#v\n", xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
	fmt.Println("--------------------")

	xi = append(xi, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	fmt.Printf("xi - %#v\n", xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
	fmt.Println("--------------------")

	xi = append(xi, 11, 12, 13)

	fmt.Printf("xi - %#v\n", xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi)) // cap is increased by 10. (not 3) (see make)
	fmt.Println("--------------------")

}
