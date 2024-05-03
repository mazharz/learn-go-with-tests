package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Maz", "")
		want := "Hello Maz!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, world!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello world!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in Spanish", func(t *testing.T) {
		got := Hello("Danny", "Spanish")
		want := "Hola Danny!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in French", func(t *testing.T) {
		got := Hello("Emily", "French")
		want := "Bonjour Emily!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in German", func(t *testing.T) {
		got := Hello("Till", "German")
		want := "Halo Till!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
