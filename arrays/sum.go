package arrays

const length = 5

// Sum takes array of int and return sum of it members
func Sum(numbers [length]int) (result int) {
	result = 0
	for i := 0; i < length; i++ {
		result += numbers[i]
	}
	return
}
