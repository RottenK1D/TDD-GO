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

	t.Run("take slices/arrays store in new slice", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{2, 8})
		want := []int{3, 10}

		// cant use != qaulity operator with slices , must use '!reflect.DeepEqual'
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("collect the total of 'tails' of each slice", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{2, 1})
		want := []int{2, 1}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
