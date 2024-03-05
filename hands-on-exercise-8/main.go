package main

import "fmt"

const (
	_ = iota // c0 == 0
	a
	b
	c
	d
	e
	f
	printStatement = "%d \t %b\n"
)

func main() {
	fmt.Printf(printStatement, 1, 1)
	fmt.Printf(printStatement, 1<<a, 1<<a)
	fmt.Printf(printStatement, 1<<b, 1<<b)
	fmt.Printf(printStatement, 1<<c, 1<<c)
	fmt.Printf(printStatement, 1<<d, 1<<d)
	fmt.Printf(printStatement, 1<<e, 1<<e)
	fmt.Printf(printStatement, 1<<f, 1<<f)
}
