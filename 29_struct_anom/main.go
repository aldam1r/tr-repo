package main

import "fmt"

func main() {
	p1 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first:     "James",
		friends:   map[string]int{"Peter": 4},
		favDrinks: []string{"Beer", "Whisky"},
	}

	p2 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first:     "M",
		friends:   map[string]int{"Astrid": 7},
		favDrinks: []string{"Red Wine", "Pina Colada"},
	}

	fmt.Println(p1)
	fmt.Println(p2)

	for k, v := range p1.friends {
		fmt.Println(p1.first, "- friends -", k, "age", v)
	}

	for _, v := range p1.favDrinks {
		fmt.Println(p1.first, "- favorite drinks -", v)
	}
}
