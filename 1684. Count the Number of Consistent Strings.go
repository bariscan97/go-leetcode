func countConsistentStrings(allowed string, words []string) int {
	res := 0
	for _, word := range words {
		var flag bool = true
		for _, char := range word {
			if !strings.Contains(allowed, string(char)) {
				flag = false
				break
			}
		}
		if flag {
			res += 1
		}
	}
	return res
}