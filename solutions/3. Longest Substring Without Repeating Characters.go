func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func lengthOfLongestSubstring(s string) int {
	dic := make(map[rune]bool)
	var res, j int
	for i, v := range s {
		for dic[v] && i > j {
			dic[rune(s[j])] = false
			j++
		}
		dic[v] = true
		res = Max(res, (i-j)+1)
	}
	return res
}