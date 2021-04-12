package diffsquares

func SquareOfSum(input int) int {
	sum := 0
	for i := 1; i <= input; i++ {
		sum += i
	}
	return sum * sum
}

func SumOfSquares(input int) int {
	sum := 0
	for i := 1; i <= input; i++ {
		sum += i * i
	}
	return sum
}

func Difference(input int) int {
	return SquareOfSum(input) - SumOfSquares(input)
}
