package arraysandslices 

func Sum(numbers [5]int) int {
	count := 0
	for _, number := range numbers {
		count += number
	}
	return count

}
