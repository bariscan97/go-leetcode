func lengthOfLongestSubstring(s string) int {
	_map := [128]bool{}
	length, max := 0, 0
	for i, j := 0, 0; i < len(s); i++ {
		index := s[i]
		if _map[index] {
			for ; _map[index]; j++ {
				length--
				_map[s[j]] = false
			}
		}

		_map[index] = true
		length++
		if length > max {
			max = length
		}
	}
	return max
}