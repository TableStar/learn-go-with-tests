package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMsg(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMsg(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMsg(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Nouvelle", "French")
		want := "Bonjour, Nouvelle"
		assertCorrectMsg(t, got, want)
	})
	t.Run("in German", func(t *testing.T) {
		got := Hello("Frederik", "German")
		want := "Hallo, Frederik"
		assertCorrectMsg(t, got, want)
	})
}

func assertCorrectMsg(t testing.TB, got, want string) {
	t.Helper() // tells the testing framework that this func is helper function
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
