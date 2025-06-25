func longestPalindrome(s string) string {
	len := len(s)
	if len == 0 {
		return ""
	}

	var l, r, pl, pr int
	for r < len {
		for r+1 < len && s[l] == s[r+1] {
			r++
		}
		for l-1 >= 0 && r+1 < len && s[l-1] == s[r+1] {
			l--
			r++
		}
		if r-l > pr-pl {
			pl, pr = l, r
		}
		l = (l+r)/2 + 1
		r = l
	}
	return s[pl : pr+1]
}