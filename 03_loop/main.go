package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
		action()
	}
}

func action() {
	x := rand.Intn(10)
	y := rand.Intn(10)
	fmt.Printf("x has the value of %v\n", x)
	fmt.Printf("y has the value of %v\n", y)

	switch {
	case x < 4 && y < 4:
		fmt.Printf("both x(%v) and y(%v) are less than 4\n", x, y)
	case x > 6 && y > 6:
		fmt.Printf("both x(%v) and y(%v) are greater than 6\n", x, y)
	case x >= 4 && x <= 6:
		fmt.Printf("x(%v) is greater or equal to 4 and less than or equal to 6\n", x)
	case y != 5:
		fmt.Printf("y(%v) is not equal to 5\n", y)
	default:
		fmt.Printf("x(%v) and y(%v) do not match any condition\n", x, y)
	}

}
