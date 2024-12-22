package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("grees", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "Chris")
		assertEqual(t, buffer.String(), "Hello, Chris")
	})
}

func assertEqual[V comparable](t testing.TB, got, want V) {
	t.Helper()
	if got != want {
		t.Errorf("got %v but wanted %v", got, want)
	}
}
