package dog

import (
	"testing"
)

func TestYears(t *testing.T) {
	got := Years(2)
	want := 20
	if got != want {
		t.Error("got", got, "want", want)
	}
}

// go test -coverprofile c.out
// go test -bench .
func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}
