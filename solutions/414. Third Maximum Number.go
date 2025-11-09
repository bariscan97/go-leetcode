func thirdMax(nums []int) int {
	sort.Ints(nums)
	N := len(nums)
	cnt := 1
	for i := N - 1; i > 0; i-- {
		if nums[i] != nums[i-1] {
			cnt++
		}
		if cnt == 3 {
			return nums[i-1]
		}
	}

	return nums[N-1]
}