func minOperations(nums []int, k int) int {
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) // 2
		if nums[mid] < k {
			l = mid + 1
		} else {
			r = mid - 1
		}

	}
	return l
}