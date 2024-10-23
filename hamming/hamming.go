package hamming

import "errors"

func Distance(a, b string) (int, error) {
	runeA := []rune(a)
	runeB := []rune(b)
	if len(runeA) != len(runeB) {
		return 0, errors.New("invalid DNA")
	}
	count := 0
	for i, v := range runeA {
		if v != runeB[i] {
			count++
		}
	}
	return count, nil
}
