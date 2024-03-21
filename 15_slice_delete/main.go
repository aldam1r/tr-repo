package main

import "fmt"

func main() {
	//Delete from a slice. (permanent)
	//slice
	xi := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("xi - %#v\n", xi)
	fmt.Println("--------------------")

	//Create new slice with the slice elements you want to keep. (There seems to be not delete ???)
	xi = append(xi[:4], xi[5:]...)
	fmt.Printf("xi - %#v\n", xi)
	fmt.Println("--------------------")

}
