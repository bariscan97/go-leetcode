func firstUniqChar(s string) int {
	_map := make(map[rune]int)
	for _, i := range s {
		_map[i]++
	}
	for idx, char := range s {
		if _map[char] == 1 {
			return idx
		}
	}
	return -1
}