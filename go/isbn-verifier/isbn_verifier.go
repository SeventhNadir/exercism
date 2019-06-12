package isbn

import (
	"errors"
	"strings"
)

//IsValidISBN will return a bool indicating whether or not the ISBN is valid or not.
//The ISBN-10 verification process is used to validate book identification numbers
//For further information please see https://en.wikipedia.org/wiki/International_Standard_Book_Number#ISBN-10_check_digit_calculation
func IsValidISBN(ISBN string) bool {
	ISBN, err := validateString(ISBN)
	if err != nil {
		return false
	}
	total := 0
	scorecard := createMap()

	for i, val := range ISBN {
		valueISBN := scorecard[string(val)]
		multiplier := 10 - i
		total += valueISBN * multiplier
	}
	if total%11 == 0 {
		return true
	}
	return false
}

func validateString(ISBN string) (string, error) {
	ISBN = strings.Replace(ISBN, "-", "", -1)
	if len(ISBN) != 10 {
		return "", errors.New("Valid ISBN numbers should be exactly 10 digits long.")
	}
	scorecard := createMap()
	for i, val := range ISBN {
		if _, ok := scorecard[string(val)]; !ok {
			return "", errors.New("This ISBN number contains an illegal character.")
		}
		if scorecard[string(val)] == 10 && i != 9 {
			return "", errors.New("This ISBN number uses X in an illegal position")
		}
	}
	return ISBN, nil
}

func createMap() map[string]int {
	scorecard := map[string]int{
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"X": 10,
	}

	return scorecard
}
