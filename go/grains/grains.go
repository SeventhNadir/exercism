package grains

import (
	"fmt"
)

//Square calculates the number of grains on a given square of a chessboard.
func Square(n int) (uint64, error) {
	if n <= 0 || n > 64 {
		return 0, fmt.Errorf("Out of acceptable range")
	}
	return 1 << uint64(n-1), nil
}

//Total calculates the number of grains on a chessboard.
func Total() uint64 {
	return 1<<64 - 1

}
