package arraysandslices 

func Sum(numbers [5]int) int {
	count := 0
	for i := 0; i < 5 ; i++ {
		count += numbers[i]
	}
	return count

}
