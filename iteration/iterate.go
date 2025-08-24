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
