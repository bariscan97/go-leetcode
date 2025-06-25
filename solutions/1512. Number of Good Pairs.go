func numIdenticalPairs(nums []int) int {
	my_map := make(map[int]int)
	res := 0
	for _, num := range nums {
		my_map[num] += 1
	}
	for _, i := range my_map {
		res += (i * (i - 1))
	}
	return int(res / 2)

}