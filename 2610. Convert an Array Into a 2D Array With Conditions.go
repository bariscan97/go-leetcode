func findMatrix(nums []int) [][]int {
	res := make([][]int, 0)
	_map := make(map[int]int)
	for _, num := range nums {
		_map[num]++
	}
	for {
		arr := make([]int, 0)
		for key, _ := range _map {
			arr = append(arr, key)
			_map[key]--
			if _map[key] == 0 {
				delete(_map, key)
			}
		}
		res = append(res, arr)
		if len(_map) == 0 {
			break
		}
	}
	return res
}