package diffsquares

func SquareOfSum(n int) int {
	return (n * (n + 1) / 2) * (n * (n + 1) / 2)
}

func SumOfSquares(n int) int {
	res := 0
	for n > 0 {
		res += n * n
		n--
	}
	return res
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
