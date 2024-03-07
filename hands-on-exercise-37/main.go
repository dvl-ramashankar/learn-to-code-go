package main

import "fmt"

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	for k, v := range m {
		fmt.Printf("key %v \t value %v\n", k, v)
	}

	fmt.Println("==================")

	age1 := m["James"]
	fmt.Println("The age of Bond :", age1)

	age2 := m["Q"]
	fmt.Println("That age of Q :", age2)

	if v, ok := m["James"]; ok {
		fmt.Println("there is james key in map. So, bond's age is :", v)
	}

	if v, ok := m["Q"]; !ok {
		fmt.Println("there is no Q key in map. So, here is the zero value of an int :", v)
	}

}
