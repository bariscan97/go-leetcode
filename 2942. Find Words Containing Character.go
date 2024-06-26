func findWordsContaining(words []string, x byte) []int {
	res := make([]int, 0)
	for idx, word := range words {
		if strings.Contains(word, string(x)) {
			res = append(res, idx)
		}
	}
	return res
}