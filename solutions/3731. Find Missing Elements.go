func findMissingElements(nums []int) []int {
	var (
		minVal int          = 101
		maxVal int          = -1
		seens  map[int]bool = make(map[int]bool)
		res    []int        = []int{}
	)

	for _, num := range nums {
		seens[num] = true
		if num > maxVal {
			maxVal = num
		}
		if num < minVal {
			minVal = num
		}
	}

	for i := minVal; i < maxVal; i++ {
		if !seens[i] {
			res = append(res, i)
		}
	}

	return res
}