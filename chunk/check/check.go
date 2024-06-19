package check

func Chunk(slice []int, size int) [][]int {
	var res [][]int
	if size <= 0{
		return nil
	}
	for i := 0; i <len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		res = append(res, slice[i:end])
	}
	return res
}


