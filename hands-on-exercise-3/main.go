package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println("Add :", add(2, 3))
	fmt.Println(split(17))
}
