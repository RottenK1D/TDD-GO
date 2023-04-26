package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "just test"}
	got := Search(dictionary, "test")
	want := "just test"

	assertString(t, got, want)
}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q ", got, want)
	}
}
