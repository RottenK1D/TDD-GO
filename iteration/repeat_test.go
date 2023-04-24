package iteration

import "testing"

func TestReapeat(t *testing.T) {
	got := Repeat("a")
	want := "aaaaa"
	assert(t, got, want)
}

func assert(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", want, got)
	}
}
