package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("printing message to person", func(t *testing.T) {
		got := Message("DUDE")
		want := "Hi, DUDE"

		if got != want {
			t.Errorf("got %q want %ql", got, want)
		}
	})

	t.Run("print 'Hi, World' if emty string is provided", func(t *testing.T) {
		got := Message("")
		want := "Hi, World"

		if got != want {
			t.Errorf("got %q want %ql", got, want)
		}
	})
}
