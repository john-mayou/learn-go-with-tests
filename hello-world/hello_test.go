package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("returns greeting to passed name", func(t *testing.T) {
		assertMessage(t, Hello("John"), "Hello, John")
	})
	t.Run("defaults to 'Hello, World' if empty name is passed", func(t *testing.T) {
		assertMessage(t, Hello(""), "Hello, World")
	})
}

func assertMessage(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
