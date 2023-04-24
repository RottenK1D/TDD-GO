package main

import "testing"

func TestHello(t *testing.T) {
	got := Message("DUDE")
	want := "Hi, DUDE"

	if got != want {
		t.Errorf("got %q want %ql", got, want)
	}
}
