func countSegments(s string) int {
	word, res := "", 0
	for _, i := range s {
		if string(i) != " " {
			word += string(i)
		} else {
			if len(word) > 0 {
				res += 1
			}
			word = ""
		}
	}
	if len(word) > 0 {
		res += 1
	}
	return res
}