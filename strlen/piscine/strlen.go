package piscine

func StrLen(s string) int {
	// l := len(s)
	runes := []rune(s)
	count := 0
	for range runes {
		count++
	}

	return count
}
