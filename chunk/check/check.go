package check

func Chunk(slice []int, size int) [][]int {
	var res [][]int
	if size <= 0 {
		return nil
	}
	for i := 0; i < len(slice); i += size {
		stp := i + size
		if stp > len(slice) {
			stp = len(slice)
		}
		res = append(res, slice[i:stp])
	}

	return res
}
