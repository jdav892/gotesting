package reflection

func walk(x interface{}, fn func(input string)) {
	fn("Lebron is the goat")
}
