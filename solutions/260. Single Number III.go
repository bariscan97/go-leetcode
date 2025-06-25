func singleNumber(nums []int) []int {
	_map := make(map[int]int)
	res := []int{}
	for _, num := range nums {
		_map[num]++
	}
	for key, val := range _map {
		if val == 1 {
			res = append(res, key)
		}
	}
	return res
}