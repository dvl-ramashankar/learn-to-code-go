package main

import "fmt"

func main() {
	c := make(chan int, 1)

	go func() {
		c <- 12
	}()

	fmt.Println(<-c)
	fmt.Printf("c:%T", c)
}
