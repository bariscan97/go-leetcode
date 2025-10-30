func minNumberOperations(target []int) int {
	res := 0
	for i := 0; i < len(target); i++ {
		if i == 0 {
			res += target[i]
		} else if target[i] > target[i-1] {
			res += target[i] - target[i-1]
		}
	}
	return res
}