package main

import (
	"fmt"
)

type canine struct {
	name string
	age  int
}

func main() {
	d := canine{
		name: "Jimmy",
		age:  dog.Years(2),
	}
	fmt.Println(d)
}
