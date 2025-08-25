package arraysandslices 

func Sum(numbers []int) int {
	count := 0
	for _, number := range numbers {
		count += number
	}
	return count

}
