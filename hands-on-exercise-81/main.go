package main

import (
	"fmt"
	"sort"
)

func main() {
	i := []int{0, 1, 6, 2, 7, 3, 8, 4, 9, 5}
	s := []string{"a", "f", "b", "g", "c", "h", "d", "i", "e"}
	fmt.Printf("Before sorting %T \t %v\n", i, i)
	sort.Ints(i)
	fmt.Printf("After sorting %T \t %v\n", i, i)
	fmt.Printf("Before sorting %T \t %v\n", s, s)
	sort.Strings(s)
	fmt.Printf("After sorting %T \t %v\n", s, s)
}
