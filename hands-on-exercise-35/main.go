package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	for k, v := range a {
		fmt.Printf("index %v, value %v \n", k, v)
	}
}
