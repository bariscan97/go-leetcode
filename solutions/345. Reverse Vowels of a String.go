func isVowel(i string) bool {
	switch i {
	case "a", "e", "i", "o", "u":
		return true
	}
	return false
}

func reverseVowels(s string) string {
	l, r := 0, len(s)-1
	arr := []rune(s)
	N := len(s)
	for {
		for !isVowel(strings.ToLower(string(arr[l]))) && l < N-1 {
			l++
		}

		for !isVowel(strings.ToLower(string(arr[r]))) && r > 0 {
			r--
		}

		if r <= l {
			break
		}

		arr[l], arr[r] = arr[r], arr[l]
		r--
		l++

	}

	return string(arr)
}