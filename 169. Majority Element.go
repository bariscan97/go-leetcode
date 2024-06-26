func majorityElement(nums []int) int {
	_map := make(map[int]int)
	for _, i := range nums {
		_map[i]++
	}
	for key, val := range _map {
		if val > len(nums)/2 {
			return key
		}
	}
	panic("hahahha")
}