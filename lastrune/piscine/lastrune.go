package piscine

func LastRune(s string) rune {
	runes := []rune(s)
	last := runes[len(runes)-1]
	return last
}
