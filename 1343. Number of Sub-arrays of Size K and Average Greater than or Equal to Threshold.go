func numOfSubarrays(arr []int, k int, threshold int) int {
	res, total, cnt := 0, 0, 0
	for idx, i := range arr {
		total += i
		cnt++
		if cnt == k {
			if total/k >= threshold {
				res += 1
			}
			total -= arr[idx-(k-1)]
			cnt--
		}
	}
	return res
}