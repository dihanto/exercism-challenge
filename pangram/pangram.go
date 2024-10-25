package pangram

import "strings"

func IsPangram(input string) bool {
	letterMap := map[rune]bool{
		'a': false, 'b': false, 'c': false, 'd': false, 'e': false, 'f': false,
		'g': false, 'h': false, 'i': false, 'j': false, 'k': false, 'l': false,
		'm': false, 'n': false, 'o': false, 'p': false, 'q': false, 'r': false,
		's': false, 't': false, 'u': false, 'v': false, 'w': false, 'x': false,
		'y': false, 'z': false,
	}

	for _, v := range strings.ToLower(input) {
		if _, exist := letterMap[v]; exist {
			letterMap[v] = true
		}
	}

	for _, v := range letterMap {
		if !v {
			return false
		}
	}
	return true
}
