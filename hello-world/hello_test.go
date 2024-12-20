package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("returns greeting to passed name", func(t *testing.T) {
		got := Hello("John")
		want := "Hello, John"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("defaults to 'Hello, World' if empty name is passed", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
