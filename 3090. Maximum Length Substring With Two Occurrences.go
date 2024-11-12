func maximumLengthSubstring(s string) int {
	dic := make(map[rune]int)
	var left int
	var res int
	for right, char := range s {
		dic[char]++
		for dic[char] > 2 {
			dic[rune(s[left])]--
			left++
		}
		diff := right - left + 1
		if diff > res {
			res = diff
		}
	}
	return res
}