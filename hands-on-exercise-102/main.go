package main

import (
	word "exercises/hands-on-exercise-102/count"
	"exercises/hands-on-exercise-102/quote"
	"fmt"
)

func main() {
	fmt.Println(word.Count(quote.QuoteText))
}
