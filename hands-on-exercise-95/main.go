package main

import "fmt"

func main() {
	c := make(chan int)
	for x := 0; x <= 10; x++ {
		go func() {
			for i := 0; i <= 10; i++ {
				c <- i
			}
		}()
	}

	for y := 0; y <= 100; y++ {
		fmt.Println(y, <-c)
	}
}
