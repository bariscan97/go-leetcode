func mergeAlternately(word1 string, word2 string) string {
	res := ""
	idx1, idx2 := 0, 0
	for {
		if idx1 < len(word1) {
			res += string(word1[idx1])
			idx1++
		}
		if idx2 < len(word2) {
			res += string(word2[idx2])
			idx2++
		}
		if idx1 == len(word1) && idx2 == len(word2) {
			break
		}
	}
	return res
}