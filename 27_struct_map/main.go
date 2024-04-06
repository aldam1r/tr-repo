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

	for _, p := range p2.ice {
		fmt.Println(p)
	}

	ms := make(map[string]person)
	ms[p1.last] = p1
	ms[p2.last] = p2
	ms[p3.last] = p3

	for k, v := range ms {
		fmt.Printf("%v \t %T \t %#v\n", k, v, v)
		for _, x := range v.ice {
			fmt.Printf("%v\n", x)
		}
	}
}
