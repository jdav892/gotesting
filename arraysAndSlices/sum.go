package arraysandslices 

func Sum(numbers []int) int {
	count := 0
	for _, number := range numbers {
		count += number
	}
	return count
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(tailsToSum ...[]int) []int {
	var sums []int
	for _, numbers := range tailsToSum {
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}
	return sums
}
