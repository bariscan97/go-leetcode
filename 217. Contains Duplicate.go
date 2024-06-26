func containsDuplicate(nums []int) bool {
	_map := make(map[int]int)
	for _, i := range nums {
		if _map[i] > 0 {
			return true
		}
		_map[i]++
	}
	return false
}