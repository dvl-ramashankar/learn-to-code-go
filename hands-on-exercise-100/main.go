package main

import (
	"exercises/hands-on-exercise-100/dog"
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
