package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to someone", func(t *testing.T) {
		got := helloWord("Dyonatha", "")
		want := "Hello, Dyonatha!"

		asserCorretMessage(t, got, want)
	})
	t.Run("Say Hello, Word! when an empty string is supplied", func(t *testing.T) {
		got := helloWord("", "")
		want := "Hello, World!"

		asserCorretMessage(t, got, want)
	})
	t.Run("In spanish", func(t *testing.T) {
		got := helloWord("Dyonatha", "Spanish")
		want := "Hola, Dyonatha!"
		asserCorretMessage(t, got, want)

	})
	t.Run("In french", func(t *testing.T) {
		got := helloWord("Dyonatha", "French")
		want := "Bonjour, Dyonatha!"
		asserCorretMessage(t, got, want)

	})
	t.Run("In Latin", func(t *testing.T) {
		got := helloWord("Dyonatha", "Latin")
		want := "Salve, Dyonatha!"
		asserCorretMessage(t, got, want)

	})
}

func asserCorretMessage(t testing.TB, got, want string) {
	// t.Helper() is needed to tell the test suite that this method is a helper.
	//By doing this when it fails the line number reported will be in our function call
	//rather than inside our test helper. This will help other developers track down problems easier.
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
