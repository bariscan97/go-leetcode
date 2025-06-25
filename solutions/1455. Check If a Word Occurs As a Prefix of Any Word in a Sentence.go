func isPrefixOfWord(sentence string, searchWord string) int {

	for index, word := range strings.Split(sentence, " ") {
		if isSamePrefix(word, searchWord) {
			return index + 1
		}
	}

	return -1
}
func isSamePrefix(w string, s string) bool {

	if len(w) < len(s) {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] != w[i] {
			return false
		}
	}
	return true
}