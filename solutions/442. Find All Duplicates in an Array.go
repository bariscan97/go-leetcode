func findDuplicates(nums []int) []int {

	_map, res := make(map[int]int), []int{}

	for _, num := range nums {
		if _map[num] > 0 {
			res = append(res, num)
		}
		_map[num]++
	}
	return res
}