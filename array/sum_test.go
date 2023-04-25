package array

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sum of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})

	t.Run("sum of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("take slices/arrays store in new slice", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{2, 8})
		want := []int{3, 10}

		// cant use != qaulity operator with slices , must use '!reflect.DeepEqual'
		checkSums(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("sum of 'tails' of each slice", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{2, 1})
		want := []int{2, 1}

		checkSums(t, got, want)
	})

	t.Run("sum safely empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}

func checkSums(t testing.TB, got, want []int) {
	t.Helper()
	// cant use != qaulity operator with slices , must use '!reflect.DeepEqual'
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
