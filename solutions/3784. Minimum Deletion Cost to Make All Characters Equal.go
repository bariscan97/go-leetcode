func minCost(s string, cost []int) int64 {
	n := len(s)

	var costSum int64
	for _, c := range cost {
		costSum += int64(c)
	}

	dic := make(map[byte]int64, n)
	for i := 0; i < n; i++ {
		dic[s[i]] += int64(cost[i])
	}

	var res int64 = 1<<63 - 1 
	for _, v := range dic {
		cur := costSum - v
		if cur < res {
			res = cur
		}
	}

	return res
}
