package reduce

func ReduceInt(a []int, f func(int, int) int) int {
	res := a[0]
	for i := 1; i < len(a); i++ {
		res = f(res, a[i])
	}
	return res
}
