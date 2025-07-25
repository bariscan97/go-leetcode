func maxSum(nums []int) int {
	dic := make(map[int]int)
	for _, n := range nums {
		dic[n]++
	}

	total := 0
	for k := range dic {
		if k > 0 {
			total += k
		}
	}

	if len(dic) > 0 && total == 0 {
		max := nums[0]
		for _, n := range nums {
			if n > max {
				max = n
			}
		}
		return max
	}

	return total
}