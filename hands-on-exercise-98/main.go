package main

import "fmt"

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("custom error: %s", ce.info)
}

func main() {
	c := customErr{
		info: "custom error",
	}

	foo(c)
}

func foo(e error) {
	fmt.Println("foo ran -", e, "\n", e.(customErr).info)
}
