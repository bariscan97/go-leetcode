func doesAliceWin(s string) bool {
	for _, char := range s {
		switch char {
		case 'a', 'e', 'i', 'o', 'u':
			return true
		}
	}
	return false
}