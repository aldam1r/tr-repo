package main

import "fmt"

// composite data type. struct
type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   42,
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			first: "Jenny",
			last:  "Moneypenny",
			age:   38,
		},
		ltk: false,
	}

	fmt.Println(sa1)
	fmt.Println(sa2)

	fmt.Printf("%T \t %#v\n", sa1, sa1)
	fmt.Printf("%T \t %#v\n", sa2, sa2)

	// both print the first name. the second is inherited.
	// When there is a "collistion" the most direct/correct variable will be used.
	fmt.Println(sa1.person.first, sa1.first)

}
