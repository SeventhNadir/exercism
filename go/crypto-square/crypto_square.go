package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

type rectangle struct {
	rows    int
	columns int
}

func Encode(plainText string) string {
	encodedString := []string{}
	normalisedString := normalise(plainText)
	rect := sizeRectangle(normalisedString)

	for i := 0; i < rect.columns; i++ {
		subString := ""
		for index, r := range normalisedString {
			if index%rect.columns == i {
				subString += string(r)
			}		
		}
		if i > 0 {
			if len(encodedString[0]) > len(subString) {
				subString += " "
			}
		}
		encodedString = append(encodedString, subString)
	}

	return strings.Join(encodedString, " ")
}

func normalise(plainText string) string {
	normalisedText := ""
	for _, r := range plainText {
		if unicode.IsDigit(r) {
			normalisedText += string(r)
		}
		if unicode.IsLetter(r) {
			r = unicode.ToLower(r)
			normalisedText += string(r)
		}
	}
	return normalisedText
}

func sizeRectangle(normalisedText string) rectangle {
	var rect rectangle
	textLength := float64(len(normalisedText))
	sqrtText := math.Sqrt(textLength)
	sqrtFloor := math.Floor(sqrtText)

	rect.rows = int(sqrtFloor)
	rect.columns = int(sqrtFloor)
	if sqrtText != sqrtFloor {
		rect.columns = int(sqrtFloor + 1)
	}
	return rect
}
