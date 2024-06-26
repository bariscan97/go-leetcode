func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s) > len(t) {
		return false
	}
	idx := 0
	for _, word := range t {
		if idx == len(s) {
			return false
		}
		if string(s[idx]) == string(word) {
			if idx == len(s)-1 {
				return true
			}
			idx++
		}
	}
	return false
}