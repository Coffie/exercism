// Package diffsquares calculates the square of sum and sum of squares for a natural number n, and its difference.
package diffsquares

// SquareOfSum calculates the square of sum for all natural numbers from 1 to input, n.
func SquareOfSum(n int) (sum int) {
	sum = (n * (n + 1))/2
	return sum*sum
}

// SumOfSquares calculates the sum of squares for all natural numbers from 1 to input, n.
func SumOfSquares(n int) (sum int) {
	sum = (n * (n + 1) * (2 * n + 1)) / 6
	return sum
}

// Difference returns the difference between the square of sum and sum of squares.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}