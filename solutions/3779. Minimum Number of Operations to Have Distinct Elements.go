func minOperations(nums []int) int {
    seen := make(map[int]bool, len(nums))
	n := len(nums)

	for i := n - 1; i >= 0; i-- {
		if seen[nums[i]] {
			idx := i + 1
			a := idx / 3
			b := idx % 3
			if b != 0 {
				a++
			}
			return a
		}
		seen[nums[i]] = true
	}

	return 0
}