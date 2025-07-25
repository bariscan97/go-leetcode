func countGoodSubstrings(s string) int {
	dic := make(map[rune]int)
	var res int
	for i, v := range s {
		dic[v]++
		if i > 1 {
			if len(dic) == 3 {
				res++
			}
			runed := rune(s[i-2])
			dic[runed]--
			if dic[runed] == 0 {
				delete(dic, runed)
			}
		}
	}
	return res
}