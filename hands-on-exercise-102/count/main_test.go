package word

import (
	"exercises/hands-on-exercise-102/quote"
	"fmt"
	"testing"
)

func TestUseCount(t *testing.T) {
	m := UseCount("one two three three three")
	for k, v := range m {
		switch k {
		case "one", "two":
			if v != 1 {
				t.Error("got", v, "want", 1)
			}
		case "three":
			if v != 3 {
				t.Error("got", v, "want", 3)
			}
		}
	}
}

func TestCount(t *testing.T) {
	n := Count("one two three")
	if n != 3 {
		t.Error("got", n, "want", 3)
	}
}

func ExampleCount() {
	fmt.Println(Count("one two three"))
	// Output:
	// 3
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.QuoteText)
	}
}

// go test -coverprofile c.out
// go tool cover -html c.out
// go test -bench .

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.QuoteText)
	}
}
