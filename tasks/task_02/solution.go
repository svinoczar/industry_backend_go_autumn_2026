package main

func rotateRunes(s string, shift int) string {
	if len(s) == 0 {
		return ""
	}

	runes := []rune(s)
	shift %= len(runes)

	if shift < 0 {
		shift += len(runes)
	}

	return string(append(runes[shift:], runes[:shift]...))
}
