package main

import (
	"fmt"
	"math/rand"
)

const printStatement = "iteration %v : x is %v \n"

func main() {
	for i := 0; i < 42; i++ {
		x := rand.Intn(5)

		switch x {
		case 0:
			fmt.Printf(printStatement, i, x)
		case 1:
			fmt.Printf(printStatement, i, x)
		case 2:
			fmt.Printf(printStatement, i, x)
		case 3:
			fmt.Printf(printStatement, i, x)
		case 4:
			fmt.Printf(printStatement, i, x)
		}
	}
}
