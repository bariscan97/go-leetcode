func minIncrementForUnique(nums []int) int {
	_map := map[int]int{}

	for _, n := range nums {
		_map[n]++
	}

	ans := 0
	curr := 0
	for len(_map) > 0 {
		if _map[curr] == 0 {
			delete(_map, curr)
			curr++
			continue
		}
		ans += _map[curr] - 1
		_map[curr+1] += _map[curr] - 1
		delete(_map, curr)
		curr++
	}

	return ans
}