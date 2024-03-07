package main

import (
	"fmt"
	"math/rand"
)

func main() {
	c := 1
	for i := 0; i < 100; i++ {
		if x := rand.Intn(5); x == 3 {
			fmt.Printf("iteration %v \t total count %v \t x is %v\n", i, c, x)
			c++
		}
	}
}
