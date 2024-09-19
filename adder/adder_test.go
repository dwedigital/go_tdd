package adder

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestAdder(t *testing.T) {

	sum := add(2, 2)
	expected := 4

	assertCorrectResult(t, sum, expected)
}

func BenchmarkAdder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(2, rand.Intn(100))
	}
}

func assertCorrectResult(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func Exampleadd() {
	sum := add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
