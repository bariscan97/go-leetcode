func largestPerimeter(nums []int) int {
	sort.Ints(nums)
	for i := len(nums) - 2; i > 0; i-- {
		if nums[i-1]+nums[i] > nums[i+1] {
			return nums[i-1] + nums[i] + nums[i+1]
		}
	}
	return 0
}