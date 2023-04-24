package main

import "testing"

func TestHello(t *testing.T) {
	got := Message()
	want := "Hi, KID"

	if got != want {
		t.Errorf("got %q want %ql", got, want)
	}
}
