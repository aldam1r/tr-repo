package main

import "fmt"

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	make  string
	model string
	doors int
	color string
}

func main() {
	v1 := vehicle{
		engine: engine{true},
		make:   "Mercedes",
		model:  "500 CL",
		doors:  4,
		color:  "Black",
	}

	v2 := vehicle{
		engine: engine{false},
		make:   "Toyota",
		model:  "Rav4",
		doors:  5,
		color:  "White",
	}

	fmt.Println(v1)
	fmt.Println(v2)

	fmt.Printf("%#v\n", v1)
	fmt.Printf("%#v\n", v2)

	fmt.Printf("%#v \t %v\n", v1.engine, v1.engine)     // seems to work but is not ok
	fmt.Printf("%#v \t %v\n", v1.electric, v1.electric) // better
	fmt.Printf("%#v \t %v\n", v2.make, v2.make)

}
