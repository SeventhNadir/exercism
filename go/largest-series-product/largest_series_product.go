package lsproduct

import (
	"errors"
	"strconv"
)

//LargestSeriesProduct will, given a string of digits, calculate the largest product for a contiguous substring of digits of length n.
func LargestSeriesProduct(digits string, span int) (int, error) {
	if span > len(digits) {
		return -1, errors.New("Slicing issue")
	}
	largestSeriesProduct := 0
	for index := 0; index <= len(digits)-span; index++ {
		subString, err := getSubstring(digits, span, index)
		if err != nil {
			return -1, err
		}
		currentSeriesProduct, err := scoreString(subString)
		if err != nil {
			return -1, err
		}
		if currentSeriesProduct > largestSeriesProduct {
			largestSeriesProduct = currentSeriesProduct
		}
	}

	return largestSeriesProduct, nil
}

func getSubstring(digits string, span int, index int) (string, error) {
	if index+span < 0 || index+span > len(digits) {
		return "", errors.New("Span and index must be non-negative and within slice.")
	}
	subString := digits[index : index+span]
	return subString, nil
}

func scoreString(subString string) (int, error) {
	score := 1
	for _, val := range subString {
		intVal, err := strconv.Atoi(string(val))
		if err != nil {
			return -1, errors.New("Can't score string containing non-integer")
		}
		score *= intVal
	}
	return score, nil
}
