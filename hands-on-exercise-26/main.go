package main

import (
	"fmt"
	"math/rand"
)

var x int

func init() {
	fmt.Println("This is from the init function")
	x = rand.Intn(400)
}

func main() {
	fmt.Printf("The value of x is %v", x)
}
