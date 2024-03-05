package main

import "fmt"

const printStatement = "%v is of type %T\n"

func main() {
	s, i, f := "Happiness", 42, 42.42
	fmt.Printf(printStatement, i, i)
	fmt.Printf(printStatement, f, f)
	fmt.Printf(printStatement, s, s)
}
