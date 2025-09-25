func kthDistinct(arr []string, k int) string {
	left := make(map[string]int)
	right := make(map[string]int)
	for _, i := range arr {
		right[i]++
	}
	for _, i := range arr {
		left[i]++
		right[i]--
		if left[i] == 1 && right[i] == 0 {
			k--
			if k == 0 {
				return i
			}
		}
	}
	return ""
}