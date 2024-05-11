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

func SumAll(numbersToSum ...[]int) []int {
	lenghtOfNumbers := len(numbersToSum)
	sums := make([]int, lenghtOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = SumSlice(numbers)
	}
	return sums
}
