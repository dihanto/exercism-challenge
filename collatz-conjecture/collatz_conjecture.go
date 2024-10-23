package collatzconjecture

import (
	"errors"
)

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("value cannot negative or zero")
	}
	count := 0
	for n != 1 {
		if n%2 == 0 {
			count++
			n = n / 2
		} else if n%2 != 0 {
			count++
			n = n*3 + 1
		}
	}

	return count, nil
}
