package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "just test"}

	t.Run("known value", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "just test"

		assertString(t, got, want)
	})

	t.Run("unknown value", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := "missing value"

		if err == nil {
			t.Fatal("expecting error")
		}
		assertString(t, err.Error(), want)
	})
}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q ", got, want)
	}
}
