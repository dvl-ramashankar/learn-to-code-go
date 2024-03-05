package main

import (
	"fmt"
)

type ByteSize int

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	printStatement = "%d \t %b\n"
)

func main() {
	fmt.Printf(printStatement, KB, KB)
	fmt.Printf(printStatement, MB, MB)
	fmt.Printf(printStatement, GB, GB)
	fmt.Printf(printStatement, TB, TB)
	fmt.Printf(printStatement, PB, PB)
	fmt.Printf(printStatement, EB, EB)
}
