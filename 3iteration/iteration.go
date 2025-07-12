package iteration

import "strings"

const repeateCount = 5

func Repeat(character string) string{
	var repeated strings.Builder

	for i := 0; i < repeateCount; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}