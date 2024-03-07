package main

import "fmt"

func main() {
	x := [5]int{}
	for i := 0; i < 5; i++ {
		x[i] = i
	}
	for k, v := range x {
		fmt.Printf("%d - %T - %d \n", v, v, k)
	}
}
