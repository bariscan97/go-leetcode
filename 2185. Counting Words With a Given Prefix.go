func prefixCount(words []string, pref string) int {
	res := 0
	for _, word := range words {
		if len(pref) > len(word) {
			continue
		}
		var flag bool = true
		for idx, char := range word {
			if idx == len(pref) {
				break
			}
			if string(pref[idx]) != string(char) {
				flag = false
				break
			}
		}
		if flag {
			res++
		}
	}
	return res
}