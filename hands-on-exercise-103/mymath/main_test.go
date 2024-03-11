package mymath

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{[]int{10, 20, 40}, 70},
		test{[]int{2, 4, 6, 8}, 20},
		test{[]int{10, 10, 10, 10, 10, 10}, 60},
	}

	for _, v := range tests {
		f := Sum(v.data)
		if f != v.answer {
			t.Error("got", f, "want", v.answer)
		}
	}

}

func ExampleSum() {
	fmt.Println(Sum([]int{1, 2, 4, 6, 9}))
	// Output:
	// 22
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10000})
	}
}
