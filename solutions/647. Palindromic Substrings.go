func countSubstrings(s string) int {
	var res int

	m := make(map[byte][]int)

	for i := 0; i < len(s); i++ {
		m[s[i]] = append(m[s[i]], i)

		if len(m[s[i]]) > 1 {
			for _, val := range m[s[i]] {
				if val != i && isPalindromic(s[val:i+1]) {
					res++
				}
			}
		}
	}

	return res + len(s)
}

func isPalindromic(s string) bool {
	if len(s) == 1 || len(s) == 0 {
		return true
	}

	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}

		l++
		r--
	}

	return true
}