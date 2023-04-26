package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "just test"}
	got := Search(dictionary, "test")
	want := "just test"

	if got != want {
		t.Errorf("got %q want %q given , %q", got, want, "test")
	}
}
