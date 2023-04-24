package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("printing message to person", func(t *testing.T) {
		got := Message("DUDE")
		want := "Hi, DUDE"
		assertCorrectMessage(t, got, want)
	})

	t.Run("print 'Hi, World' if emty string is provided", func(t *testing.T) {
		got := Message("")
		want := "Hi, World"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
