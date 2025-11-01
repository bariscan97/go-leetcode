func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	dic := make(map[rune]int)

	for _, i := range s {
		dic[i]++
	}

	size := len(dic)

	for _, i := range t {
		if dic[i] == 1 {
			size--
		}
		dic[i]--
	}

	return size == 0
}