package series

func All(n int, s string) []string {
	substrings := []string{}
	for i := 0; i <= len(s)-n; i++ {
		substrings = append(substrings, s[i:i+n])
	}
	return substrings
}

func UnsafeFirst(n int, s string) string {
	return s[0:n]
}

func First(n int, s string) (first string, ok bool) {
	if len(s) < 0 || n < 0 || n > len(s) {
		return "", false
	}
	return s[0:n], true
}