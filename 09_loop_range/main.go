package main

import "fmt"

func main() {
	xi := []int{42, 43, 44, 45, 46, 47}
	for i, v := range xi { // i, v could be a nice standard. Meaning index, value.
		fmt.Printf("Index %v, value %v\n", i, v)
	}
	for _, x := range xi {
		fmt.Printf("Value %v\n", x)
	}

	m := map[string]int{
		"James":      42,
		"Moneypenny": 38,
	}
	for n, a := range m { // k, v could be a nice standard. Meaning key, value.
		fmt.Printf("My name is %v and I am %v years old\n", n, a)
	}
	age := m["James"]
	fmt.Println(age)

	if a, ok := m["James"]; ok {
		fmt.Printf("James' age is %v\n", a)
	} else {
		fmt.Println("James' age is not in the list")
	}

	if a, ok := m["Q"]; ok {
		fmt.Printf("Q's age is %v\n", a)
	} else {
		fmt.Println("Q's age is not in the list")
	}

}
