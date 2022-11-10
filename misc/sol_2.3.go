package misc

import (
	"fmt"
	"strings"
)

func Sol2_3(word string, k int) string {
	result := []string{}
	for i, rune := range word {
		char := fmt.Sprintf("%c", rune)
		if (i+1)%k != 0 {
			result = append(result, char)
		}
	}
	return strings.Join(result, "")
}
