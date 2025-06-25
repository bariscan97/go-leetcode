func runningSum(nums []int) []int {
	prefix := 0
	res := make([]int, 0)
	for _, num := range nums {
		prefix += num
		res = append(res, prefix)
	}
	return res
}