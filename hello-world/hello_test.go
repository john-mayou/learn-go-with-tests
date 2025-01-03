package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("returns greeting to passed name", func(t *testing.T) {
		assertMessage(t, Hello("John", ""), "Hello, John")
		assertMessage(t, Hello("Jane", ""), "Hello, Jane")
		assertMessage(t, Hello("First Last", ""), "Hello, First Last")
	})
	t.Run("returns 'Hello, World' as greeting if empty name is passed", func(t *testing.T) {
		assertMessage(t, Hello("", ""), "Hello, World")
	})
	t.Run("returns greeting in different languages", func(t *testing.T) {
		assertMessage(t, Hello("John", ""), "Hello, John")
		assertMessage(t, Hello("John", "english"), "Hello, John") // lower case
		assertMessage(t, Hello("John", "English"), "Hello, John")
		assertMessage(t, Hello("John", "Spanish"), "Hola, John")
		assertMessage(t, Hello("John", "French"), "Bonjour, John")
	})
}

func assertMessage(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
