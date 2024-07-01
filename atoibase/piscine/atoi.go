package piscine

func AtoiBase(s string, base string) int {
	res := 0
	l := len(base)
	count := make(map[rune]int)
	if l < 2 {
		return 0
	}
	for _, char := range base {
		if char == '+' || char == '-' {
			return 0
		}
	}
	for i, char := range base {
		count[char] = i 
	}
	for _, char := range s {
		value, seen := count[char]
		if !seen {
			return 0
		}
		res = res*l + value 
	}
	return res
}
