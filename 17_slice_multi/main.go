package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Martini", "Chocolate"}
	mm := []string{"Marjorie", "Moneypenny", "Ice-tea", "Strawberry", "Plum"}
	fmt.Println(jb)
	fmt.Println(mm)

	xxs := [][]string{jb, mm}
	fmt.Println(xxs)
}
