package arrays

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := [length]int{1, 2, 3, 4, 5}
		result := Sum(numbers)
		expected := 15

		if result != expected {
			t.Errorf("expected '%d' but got '%d', input %v", expected, result, numbers)
		}

	})
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := SumSlice(numbers)
		expected := 6

		if result != expected {
			t.Errorf("expected '%d' but got '%d', input %v", expected, result, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9}, []int{2, 2})
	want := []int{3, 9, 4}

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTail(t *testing.T) {
	got := SumAllTail([]int{1, 2}, []int{0, 9}, []int{2, 3, 1})
	want := []int{2, 9, 4}

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
