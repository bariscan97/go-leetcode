func reverseStr(s string, k int) string {
	str := []byte(s)

	for i := 0; i < len(s); i = i + k*2 {
		reverse(str[i:min(i+k, len(s))])
	}

	return string(str)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func reverse(s []byte) {
	left := 0
	right := len(s) - 1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}