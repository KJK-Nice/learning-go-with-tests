package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	t.Run("repeat a char 5 times", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"
		assertCorrectResult(t, got, want)
	})

	t.Run("repeat a word 3 times", func(t *testing.T) {
		got := Repeat("nice", 3)
		want := "nicenicenice"
		assertCorrectResult(t, got, want)
	})
}

func assertCorrectResult(t testing.TB, repeated, expected string) {
	t.Helper()
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	result := Repeat("c", 5)
	fmt.Println(result)
	// Output: ccccc
}
