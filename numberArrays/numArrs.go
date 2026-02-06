package numberarrays

// Sum calculates the total from a slice of numbers
func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAllTails calculates the sum of all but the first number given
func SumAllStails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}

func Reduce[T any](collection []T, f func(T, T) T, initialValue T) T {
	result := initialValue
	for _, x := range collection {
		result = f(result, x)
	}
	return result
}
