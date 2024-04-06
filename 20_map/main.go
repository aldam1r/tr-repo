package main

import "fmt"

func main() {
	// map
	am := map[string]int{
		"Pieter": 60,
		"Simon":  55,
	}
	fmt.Println("Age of Pieter is", am["Pieter"])
	fmt.Println(am)
	fmt.Printf("%#v\n", am)
	fmt.Println("----------------------")

	// map: make, fill, range
	an := make(map[string]int)
	an["Lucas"] = 28
	an["Bert"] = 32

	for k, v := range an {
		fmt.Println(k, v)
	}
	for _, x := range an {
		fmt.Println(x)
	}
	fmt.Println("----------------------")

	// map: delete
	delete(an, "Bert")
	for k, v := range an {
		fmt.Println(k, v)
	}
	fmt.Println("----------------------")

	// map: non-existing. delete, print
	delete(an, "Klaas") // No panic.
	for k, v := range an {
		fmt.Println(k, v)
	}
	fmt.Println(an["Klaas"]) // prints the empty/zero value.
	fmt.Println("----------------------")

	// map: comma ok,
	v, ok := an["klaas"]
	if ok {
		fmt.Println(v)
	}

	// combined with stmt;stmt
	if w, ok := an["klaas"]; ok {
		fmt.Println(w)
	} else {
		fmt.Println("Not ok")
	}

}
