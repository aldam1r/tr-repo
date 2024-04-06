package main

import "fmt"

type person struct {
	first string
	last  string
	ice   []string
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		ice:   []string{"Chocolate", "Stawberry"},
	}

	p2 := person{
		first: "Jenny",
		last:  "Moneypenny",
		ice:   []string{"Vanila", "Peach"},
	}

	p3 := person{
		first: "Professor",
		last:  "Q",
		ice:   []string{"Pistache", "Melon", "Lemon"},
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)

	fmt.Printf("%T \t %#v\n", p1, p1)
	fmt.Printf("%T \t %#v\n", p2, p2)
	fmt.Printf("%T \t %#v\n", p3, p3)

	for _, p := range p1.ice {
		fmt.Println(p)
	}

}
