package main

import "fmt"

var c, python, java = true, false, "no"

func main() {
	i, b := 1, 2.0
	fmt.Printf("%T \t %T \t %T \t %T \t %T ", i, b, c, python, java)
}
