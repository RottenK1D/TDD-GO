package iteration

import "testing"

func TestReapeat(t *testing.T) {
	t.Run("repeat char 5 times", func(t *testing.T) {
		got := Repeat("a", 0)
		want := "aaaaa"
		assert(t, got, want)
	})

	t.Run("caller specifies repeat count", func(t *testing.T) {
		got := Repeat("a", 10)
		want := "aaaaaaaaaa"
		assert(t, got, want)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 100000)
	}
}

func assert(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", want, got)
	}
}
