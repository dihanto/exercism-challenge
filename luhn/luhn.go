package luhn

import (
	"strings"
	"unicode"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")

	if len(id) < 2 {
		return false
	}
	for _, char := range id {
		if !unicode.IsDigit(char) {
			return false
		}
	}

	sum := 0
	n := len(id)

	for i, char := range id {
		digit := int(char - '0')

		if (n-i)%2 == 0 {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}

	return sum%10 == 0
}
