func countGoodSubstrings(s string) int {
	if len(s) < 3 {
		return 0
	}
	var res int = 0
	_map := make(map[string]int)

	for _, i := range s[0:3] {
		_map[string(i)]++
	}
	if len(_map) == 3 {
		res++
	}
	var left int = 0
	for _, i := range s[3:] {
		_map[string(i)]++
		_map[string(s[left])]--
		if _map[string(s[left])] == 0 {
			delete(_map, string(s[left]))
		}
		left++
		if len(_map) == 3 {
			res++
		}
	}
	return res
}