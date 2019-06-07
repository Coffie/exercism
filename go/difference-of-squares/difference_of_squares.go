// Package diffsquares calculates the square of sum and sum of squares for a natural number n, and its difference.
package diffsquares

// SquareOfSum calculates the square of sum for all natural numbers from 1 to input, n.
func SquareOfSum(n int) int {
	return n * n * (n + 1) * (n + 1) / 4
}

// SumOfSquares calculates the sum of squares for all natural numbers from 1 to input, n.
func SumOfSquares(n int) int {
	return n * (n + 1) * (2 * n + 1) / 6
}

// Difference returns the difference between the square of sum and sum of squares.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}