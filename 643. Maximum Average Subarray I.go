func findMaxAverage(nums []int, k int) float64 {
	res := -10548298
	var total int
	for index, num := range nums {
		total += num
		if index+1 >= k {
			if res < total {
				res = total
			}
			total -= nums[index-(k-1)]
		}
	}
	return float64(res) / float64(k)
}