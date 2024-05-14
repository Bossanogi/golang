package arrays

import (
	"slices"
	"testing"
)

var input = []int{0, 1, 2, 3, 4, 5}
var input2 = []int{2, 3, 4, 6, 8, 9}

func isEven(num int) bool {
	if num == 0 {
		return false
	}
	if (num % 2) == 0 {
		return true
	} else {
		return false
	}
}

func TestFilter(t *testing.T) {
	t.Run("with isEven predicate", func(t *testing.T) {
		input := input
		got := Filter(input, isEven)
		want := []int{2, 4}

		if !(slices.Equal(got, want)) {
			t.Errorf("got %v want %v, input %v", got, want, input)
		}

	})
	t.Run("with anonymous predicate", func(t *testing.T) {
		input := input2
		got := Filter(input, func(num int) bool {
			if num == 0 {
				return false
			}
			if (num % 3) == 0 {
				return true
			} else {
				return false
			}
		})
		want := []int{3, 6, 9}

		if !(slices.Equal(got, want)) {
			t.Errorf("got %v want %v, input %v", got, want, input)
		}
	})
}
