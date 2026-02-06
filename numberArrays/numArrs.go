package numberarrays

// Sum calculates the total from a slice of numbers
func Sum(numbers []int) int {
	add := func(acc, x int) int { return acc + x }
	return Reduce(numbers, add, 0)
}

// SumAllTails calculates the sum of all but the first number given
func SumAllStails(numbers ...[]int) []int {
	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			tail := x[1:]
			return append(acc, Sum(tail))
		}
	}
	return Reduce(numbers, sumTail, []int{})
}

func Reduce[T any](collection []T, f func(T, T) T, initialValue T) T {
	result := initialValue
	for _, x := range collection {
		result = f(result, x)
	}
	return result
}
