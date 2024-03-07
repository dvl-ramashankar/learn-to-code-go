package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}

	part2()
}

func part2() {
	for i := 0; i < 100; i++ {
		x := rand.Intn(10)
		y := rand.Intn(10)
		fmt.Printf("iteration %v : x and y are %v and %v \n", i, x, y)

		switch {
		case x < 4 && y < 4:
			fmt.Println("both are less than 4")
		case x > 6 && y > 6:
			fmt.Println("both are greater than 6")
		case x >= 4 && x <= 6:
			fmt.Println("x is greater than or equal to 4 and y is less than or greater than 6")
		case y != 5:
			fmt.Println("y is not 5")
		default:
			fmt.Println("none of the previous were met")
		}
	}
}
