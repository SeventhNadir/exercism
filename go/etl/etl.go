package etl

import "strings"

//Transform takes our legacy int to array mapping and converts it to a string to int mapping.
func Transform(input map[int][]string) map[string]int {
	transformed := map[string]int{}
	for score, letters := range input {
		for _, letter := range letters {
			transformed[string(strings.ToLower(letter))] = score
		}
	}
	return transformed
}
