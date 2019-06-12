package secret

func Handshake(handshake uint) []string {
	code := []string{}
	if handshake&1 == 1 {
		code = append(code, "wink")
	}
	if handshake&2 == 2 {
		code = append(code, "double blink")
	}
	if handshake&4 == 4 {
		code = append(code, "close your eyes")
	}
	if handshake&8 == 8 {
		code = append(code, "jump")
	}
	if handshake&16 == 16 {
		code = reverse(code)
	}``
	return code
}

func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
