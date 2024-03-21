package main

import "fmt"

func main() {
	//slice a slice. (Not permanent)
	//slice
	xi := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("xi - %v\n", xi)
	fmt.Printf("xi - %#v\n", xi)
	fmt.Println("--------------------")

	//[from inclusive:to exlusive]
	fmt.Printf("xi - %v\n", xi[2:4])
	fmt.Printf("xi - %#v\n", xi[2:4])
	fmt.Println("--------------------")

	//[:to exlusive]
	fmt.Printf("xi - %v\n", xi[:6])
	fmt.Printf("xi - %#v\n", xi[:6])
	fmt.Println("--------------------")

	//[from inclusive:]
	fmt.Printf("xi - %v\n", xi[3:])
	fmt.Printf("xi - %#v\n", xi[3:])
	fmt.Println("--------------------")
}
