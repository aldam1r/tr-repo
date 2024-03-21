package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	b := a

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("--------------------")

	a[0] = 7

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("--------------------")

	c := []int{0, 1, 2, 3, 4, 5}
	d := make([]int, 6)
	copy(d, c)

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println("--------------------")

	c[0] = 7

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println("--------------------")

}
