package main

import "fmt"

var a = 1

const b = 2

func main() {
	c := 3.0
	fmt.Printf("The value of a is %v and type is %T\n", a, a)
	fmt.Printf("The value of b is %v and type is %T\n", b, b)
	fmt.Printf("The value of c is %v and type is %T\n", c, c)
}
