package iteration

func Repeat(char string) string {
	fill := "a"
	str := ""

	for i := 0; i < 5; i++ {
		str += fill
	}
	return str
}
