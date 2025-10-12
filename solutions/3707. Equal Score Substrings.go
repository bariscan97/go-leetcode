func scoreBalance(s string) bool {
	var (
		N          int = len(s)
		r          int = N - 1
		l          int = 0
		pref, suff int32
	)
	for r >= l {
		if pref < suff {
			pref += 26 - ('z' - rune(s[l]))
			l++
		} else {
			suff += 26 - ('z' - rune(s[r]))
			r--
		}
	}
	return pref == suff
}