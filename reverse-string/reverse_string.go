package reverse

func Reverse(input string) string {
	runes := []rune(input)
	n := len(runes)
	var result []rune

	for i := n - 1; i >= 0; i-- {
		result = append(result, runes[i])
	}

	return string(result)
}
