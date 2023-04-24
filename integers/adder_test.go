package integers

import "testing"

func TestAdder(t *testing.T) {
	t.Run("add 2 interers", func(t *testing.T) {
		got := Add(2, 2)
		want := 4
		assert(t, got, want)
	})
}

func assert(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
