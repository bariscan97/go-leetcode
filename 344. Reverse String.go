func reverseString(s []byte) {
	n := len(s) - 1
	for i := 0; i < int(len(s)/2); i++ {
		prev := s[i]
		s[i] = s[n-i]
		s[n-i] = prev
	}
}