package main

import "fmt"

func main() {
	// map: string, slice, loop
	ms := make(map[string][]string)
	ms["bond_james"] = []string{"shaken, not stirred", "martinis", "fast cars"}
	ms["moneypenny_jenny"] = []string{"intelligence", "literature", "computer science"}
	ms["no_dr"] = []string{"cats", "ice cream", "sunsets"}
	for k, v := range ms {
		fmt.Println("Key:", k)
		for i, x := range v {
			fmt.Printf("Value: %v \t Index: %v\n", x, i)
		}
	}
	fmt.Println("-------------------")

	// map: add
	ms["flemming_iam"] = []string{"steaks", "sigars", "espionage"}
	for k, v := range ms {
		fmt.Println("Key:", k)
		for i, x := range v {
			fmt.Printf("Value: %v \t Index: %v\n", x, i)
		}
	}

	fmt.Println("-------------------")

	// map: delete

	delete(ms, "no_dr")
	for k, v := range ms {
		fmt.Println("Key:", k)
		for i, x := range v {
			fmt.Printf("Value: %v \t Index: %v\n", x, i)
		}
	}

	fmt.Println("-------------------")

}
