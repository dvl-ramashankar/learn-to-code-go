package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)
	fmt.Printf("x and y are %v and %v \n", x, y)

	if x < 4 && y < 4 {
		fmt.Println("both are less than 4")
	} else if x > 6 && y > 6 {
		fmt.Println("both are greater than 6")
	} else if x >= 4 && x <= 6 {
		fmt.Println("x is greater than or equal to 4 and y is less than or greater than 6")
	} else if y != 5 {
		fmt.Println("y is not 5")
	} else {
		fmt.Println("none of the previous were met")
	}
}
