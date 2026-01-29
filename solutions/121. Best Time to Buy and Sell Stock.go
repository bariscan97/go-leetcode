func maxProfit(prices []int) int {
	var (
		mini int = 1_000_00_1
		res  int
	)

	for _, i := range prices {
		res = max(res, i-mini)
		mini = min(mini, i)
	}

	return res
}