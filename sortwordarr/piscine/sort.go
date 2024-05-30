package piscine

func SortWordArr(a []string) {
	l := len(a)
	for i := 0; i < l; i++ {
		for j := 0; j < l-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
