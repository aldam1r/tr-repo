package main

import "fmt"

func main() {
	// create array
	a := [5]int{}
	// assign values
	for i := 1; i <= 5; i++ {
		a[i-1] = i
	}
	// print type and value
	for i, v := range a {
		fmt.Printf("%T \t %v\t %v\n", v, v, i)
	}

	// create slice
	b := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for i, v := range b {
		fmt.Printf("%T\t%v\t%v\n", v, v, i)
	}
	fmt.Println("--------------------")

	//[inclusive:exclusive]
	fmt.Println(b[:5])
	fmt.Println(b[5:])
	fmt.Println(b[2:7])
	fmt.Println(b[1:6])
	fmt.Println("--------------------")

	// append
	c := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	c = append(c, 52)
	fmt.Println(c)
	c = append(c, 53, 54, 55)
	fmt.Println(c)
	d := []int{56, 57, 58, 59, 60}
	c = append(c, d...)
	fmt.Println(c)
	fmt.Println("--------------------")

	e := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	e = append(e[:3], e[6:]...)

	fmt.Println(e)
	fmt.Println("--------------------")

	// make
	f := make([]string, 0, 50)
	f = append(f, "Alabama", "Alaska", "Arizona", "Arkansas", "California", "Colorado", "Connecticut",
		"Delaware", "Florida", "Georgia", "Hawaii", "Idaho", "Illinois", "Indiana", "Iowa", "Kansas",
		"Kentucky", "Louisiana", "Maine", "Maryland", "Massachusetts", "Michigan", "Minnesota",
		"Mississippi", "Missouri", "Montana", "Nebraska", "Nevada", "New Hampshire", "New Jersey",
		"New Mexico", "New York", "North Carolina", "North Dakota", "Ohio", "Oklahoma", "Oregon",
		"Pennsylvania", "Rhode Island", "South Carolina", "South Dakota", "Tennessee", "Texas",
		"Utah", "Vermont", "Virginia", "Washington", "West Virginia", "Wisconsin", "Wyoming")
	for i := 0; i < len(f); i++ {
		fmt.Printf("Index: %v \t state: %v\n", i, f[i])
	}

	fmt.Println("--------------------")

	j := []string{"James", "Bond", "Shaken, not strirred"}
	m := []string{"Miss", "Moneypenny", "I'm 008"}

	xxs := [][]string{j, m}

	for _, v := range xxs {
		fmt.Println(v)
	}
	for _, v := range xxs {
		for _, x := range v {
			fmt.Println(x)
		}
	}

}
