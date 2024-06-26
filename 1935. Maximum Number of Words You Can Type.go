func canBeTypedWords(text string, brokenLetters string) int {
	words := strings.Split(text, " ")
	var res int = 0
	for _, word := range words {
		var flag bool = true
		for _, char := range brokenLetters {
			if includes(string(char), word) {
				flag = false
				break
			}
		}
		if flag {
			res += 1
		}
	}
	return res
}
func includes(val string, word string) bool {
	for _, i := range word {
		if val == string(i) {
			return true
		}
	}
	return false
}