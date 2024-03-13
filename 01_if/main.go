package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)
	fmt.Printf("x has the value of %v\n", x)
	fmt.Printf("y has the value of %v\n", y)
	if x < 4 && y < 4 {
		fmt.Printf("both x(%v) and y(%v) are less than 4\n", x, y)
	} else if x > 6 && y > 6 {
		fmt.Printf("both x(%v) and y(%v) are greater than 6\n", x, y)
	} else if x >= 4 && x <= 6 {
		fmt.Printf("x(%v) is greater or equal to 4 and less than or equal to 6\n", x)
	} else if y != 5 {
		fmt.Printf("y(%v) is not equal to 5\n", y)
	} else {
		fmt.Printf("x(%v) and y(%v) do not match any condition\n", x, y)
	}

}
