package reverse

type runeSlice []rune

func (r runeSlice) Swap(i, j int) { r[i], r[j] = r[j], r[i] }

func String(input string) string {
	if len(input) == 0 {
		return ""
	}
	var runes runeSlice = []rune(input)
	for i := 0; i <= len(runes)/2; i++ {
		runes.Swap(i, len(runes)-1-i)
	}
	return string(runes)
}
