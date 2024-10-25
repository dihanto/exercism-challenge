package isogram

import "strings"

func IsIsogram(word string) bool {
	word = strings.ReplaceAll(word, " ", "")
	word = strings.ReplaceAll(word, "-", "")
	word = strings.ToUpper(word)
	letterMap := make(map[rune]bool)
	for _, char := range word {
		if letterMap[char] {
			return false
		}
		letterMap[char] = true
	}
	return true
}
