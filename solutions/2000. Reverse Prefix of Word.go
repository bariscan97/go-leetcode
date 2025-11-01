func reverse(s string) string {
	N := len(s)
	res := []byte(s)
	for i := 0; i < N/2; i++ {
		res[i], res[N-1-i] = res[N-1-i], res[i]
	}
	return string(res)
}

func reversePrefix(word string, ch byte) string {
	for i, chr := range word {
		if chr == rune(ch) {
			return reverse(word[:i+1]) + word[i+1:]
		}
	}
	return word
}