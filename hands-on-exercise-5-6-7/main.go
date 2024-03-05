package main

import (
	"fmt"
	"math"
)

const Pi = 3.14

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	var x, y int = 3, 4
	var fl float64 = math.Sqrt(float64(x*x + y*y))
	fmt.Println(fl)
	v := 42
	fmt.Printf("v is of type %T\n", v)
	fmt.Println("Pi :", Pi)

	fmt.Printf("%d \t %b \t \n", 1, 1)
	fmt.Printf("%d \t %b \t \n", 1>>1, 1>>1)
}
