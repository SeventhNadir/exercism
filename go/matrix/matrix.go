package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Matrix is a simple array of int arrays.
type Matrix [][]int

// Cols returns a columnar view of the matrix
func (m *Matrix) Cols() [][]int {
	var nCols int
	if len(*m) > 0 {
		nCols = len((*m)[0])
	}
	matrix := make([][]int, nCols)
	for _, row := range *m {
		for i, col := range row {
			matrix[i] = append(matrix[i], col)
		}
	}
	return matrix
}

// Rows return a row based view of the matrix
func (m *Matrix) Rows() [][]int {
	matrix := make(Matrix, len(*m))
	for i, row := range *m {
		matrix[i] = make([]int, len(row))
		for j, val := range row {
			matrix[i][j] = val
		}
	}
	return matrix
}

// Set will set the value at a given row and column. true if success, else false.
func (m *Matrix) Set(row, col, value int) bool {
	if row < 0 || col < 0 || row >= len(*m) || (len(*m) > 0 && col >= len((*m)[0])) {
		return false
	}
	(*m)[row][col] = value
	return true
}

// New takes an input string and returns a matrix
func New(input string) (*Matrix, error) {
	matrix := Matrix{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		matrixMembers := strings.Fields(line)
		if len(matrix) > 0 && len(matrixMembers) != len(matrix[len(matrix)-1]) {
			return nil, errors.New("uneven rows")
		}
		m := make([]int, len(matrixMembers))
		for i, member := range matrixMembers {
			num, err := strconv.Atoi(member)
			if err != nil {
				return nil, err
			}
			m[i] = num
		}
		matrix = append(matrix, m)
	}
	return &matrix, nil
}
