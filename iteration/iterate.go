package iteration

import "strings"

const maxCount = 5


func Repeat(char string) string {
	var repeated strings.Builder
	for i := 0; i < maxCount; i++ {
		repeated.WriteString(char)
	}
	return repeated.String()
}

func Repeater(num int) int {
 	bucket := 0
	for i := 0; i < maxCount; i++ {
		bucket += num
	}
	return bucket
}
