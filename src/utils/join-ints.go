package utils

import "fmt"

func JoinInts(ints []int, delimiter string) string {
	result := ""
	for i, num := range ints {
		result += fmt.Sprintf("%d", num)
		if i < len(ints)-1 {
			result += delimiter
		}
	}
	return result
}
