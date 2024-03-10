package main

import "fmt"

func main() {
	c := send()
	recieve(c)
}

func send() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i <= 100; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func recieve(c <-chan int) {
	for k := range c {
		fmt.Println(k)
	}

}
