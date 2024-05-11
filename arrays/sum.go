package arrays

const length = 5

// Sum takes array of int and return sum of it members
func Sum(numbers [length]int) (result int) {
	result = 0
	for _, number := range numbers {
		result += number
	}
	return
}

func SumSlice(numbers []int) (result int) {
	result = 0
	for _, number := range numbers {
		result += number
	}
	return
}

func SumAll(slice1 []int, slice2 []int) (sum []int) {
	sum = nil
	sum = append(sum, SumSlice(slice1))
	sum = append(sum, SumSlice(slice2))

	return
}
