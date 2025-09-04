package iteration

import "strings"

// Repeat a string by specific number
func Repeat(charecter string, repeatCount int) string {
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(charecter)
	}
	return repeated.String()
}
