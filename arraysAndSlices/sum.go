package arraysandslices 

func Sum(numbers []int) int {
	count := 0
	for _, number := range numbers {
		count += number
	}
	return count
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}
