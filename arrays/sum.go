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

func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, SumSlice(numbers))
	}
	return
}

func SumAllTail(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {

		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, SumSlice(tail))
		}
	}

	return
}
