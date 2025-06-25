func reverseWords(s string) string {

	bytes := []byte(s)
	var start, end int
	lastSpaceIndex := -1

	for i := 0; i <= len(bytes); i++ {
		if i == len(bytes) || bytes[i] == ' ' {
			start, end = lastSpaceIndex+1, i-1
			for i, j := start, end; i < j; i, j = i+1, j-1 {
				bytes[i], bytes[j] = bytes[j], bytes[i]
			}
			lastSpaceIndex = i
		}
	}

	return string(bytes)
}