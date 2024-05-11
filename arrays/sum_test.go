package arrays

import "testing"

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
