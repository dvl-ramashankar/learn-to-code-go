package main

import "fmt"

func main() {
	i := 1
	for {
		if i > 10 {
			break
		}
		fmt.Println(i)
		i++
	}

}
