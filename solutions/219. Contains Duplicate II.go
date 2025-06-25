func containsNearbyDuplicate(nums []int, k int) bool {
	_map := make(map[int]int)
	for idx, num := range nums {
		if _map[num] > 0 && (idx+1)-_map[num] <= k {
			return true
		}
		_map[num] = idx + 1
	}
	return false
}