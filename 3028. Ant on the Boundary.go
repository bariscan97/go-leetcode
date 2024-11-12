func returnToBoundaryCount(nums []int) int {
	var (
		pos int
		res int
	)
	for _, num := range nums {
		pos += num
		if pos == 0 {
			res += 1
		}
	}
	return res
}