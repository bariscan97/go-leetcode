func singleNumber(nums []int) int {
	_map := make(map[int]int)
	for _, num := range nums {
		_map[num]++
	}
	for key, val := range _map {
		if val == 1 {
			return key
		}
	}
	panic("")
}