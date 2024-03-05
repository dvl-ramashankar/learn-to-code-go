package main

import "fmt"

const printStatement = "%d \t %b\n"

func main() {
	fmt.Printf(printStatement, 1, 1)
	fmt.Printf(printStatement, 3, 3)
	fmt.Printf(printStatement, 4, 4)
	fmt.Printf(printStatement, 2, 2)
	fmt.Printf(printStatement, 5, 5)
}
