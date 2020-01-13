package grains

import (
	"errors"
	"fmt"
	"math"
)

// Square calculates how many grains of rice a given square of a chess board
// will have when you double the amount for each square. Uncommenting squareSpeed
// uses a faster calculation for speed optimization.
func Square(in int) (uint64, error) {
	if in <= 0  || in > 64 {
		return 0, errors.New(fmt.Sprintf("Invalid number of squares given. Range 0 < input <= 64: got '%d'", in))
	}
	if in <= 2 {
		return uint64(in), nil
	}
	return uint64(math.Pow(2, float64(in - 1))), nil
	//return squareSpeed(in)
}

// squareSpeed is optimized version of Square based on speed of algorithm
// using a loop to calculate the output
func squareSpeed(in int) (uint64, error) {
	var sum uint64 = 2
	for i := 2; i <= in - 1; i++ {
		sum *= 2
	}
	return sum, nil
}

// Total returns the total amount of grains needed to fill a chess board with
// grains when doubling the amount for each square.
// todo: implement Total using channels and goroutine to check if speed improves
func Total() uint64 {
	var sum uint64
	for i := 1; i <= 64; i++ {
		tmp, err := Square(i)
		if err != nil {
			panic(err)
		}
		sum += tmp
	}
	return sum
}