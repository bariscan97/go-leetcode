func maximumHappinessSum(happiness []int, k int) int64 {
	N := len(happiness)
	sort.Ints(happiness)

	var res, cnt, val int

	for i := N - 1; N-i <= k; i-- {
		val = happiness[i] - cnt
		if val > 0 {
			res += val
		} else {
			break
		}
		cnt++
	}

	return int64(res)
}