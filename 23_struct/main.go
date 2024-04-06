package main

import "fmt"

// composite data type. struct
type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   42,
	}

	p2 := person{
		first: "Jenny",
		last:  "Moneypenny",
		age:   38,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Printf("%T \t %#v\n", p1, p1)
	fmt.Printf("%T \t %#v\n", p2, p2)

}
