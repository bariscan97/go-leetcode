func maximumWealth(accounts [][]int) int {
	res := 0
	for _, account := range accounts {
		sum := 0
		for _, i := range account {
			sum += i
		}
		res = max(res, sum)
	}
	return res
}