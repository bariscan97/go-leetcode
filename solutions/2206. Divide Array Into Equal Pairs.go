func divideArray(nums []int) bool {

	_map := make(map[int]int)
	for _, i := range nums {
		_map[i]++
	}
	for _, val := range _map {
		if val%2 > 0 {
			return false
		}
	}
	return true
}