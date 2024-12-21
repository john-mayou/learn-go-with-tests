package iteration

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	output := Repeat("a", 3)
	fmt.Println(output)
	// Output: aaa
}

func TestRepeat(t *testing.T) {
	t.Run("returns expected string", func(t *testing.T) {
		assertEqualStr(t, Repeat("a", 2), "aa")
		assertEqualStr(t, Repeat("a", 3), "aaa")
		assertEqualStr(t, Repeat("ab", 3), "ababab")
	})
}

func assertEqualStr(t testing.TB, actual, expected string) {
	t.Helper()
	if actual != expected {
		t.Errorf("expected %q but got %q", expected, actual)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for range b.N {
		Repeat("a", 3)
	}
}
