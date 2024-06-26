func findLonely(nums []int) []int {
	_map := make(map[int]int)

	for _, num := range nums {
		_map[num]++
	}

	lonely := make([]int, 0)

	for _, num := range nums {
		if _map[num] == 1 && _map[num-1] == 0 && _map[num+1] == 0 {
			lonely = append(lonely, num)
		}
	}
	return lonely
}