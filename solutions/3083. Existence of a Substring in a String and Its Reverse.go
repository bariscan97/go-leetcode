func isSubstringPresent(s string) bool {
	visited := make(map[string]bool)
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] || visited[string(s[i+1])+string(s[i])] {
			return true
		}
		visited[string(s[i])+string(s[i+1])] = true
	}
	return false
}