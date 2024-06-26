func countWords(words1 []string, words2 []string) int {
	_map1, _map2 := make(map[string]int), make(map[string]int)
	res := 0
	for _, i := range words1 {

		_map1[string(i)]++
	}
	for _, i := range words2 {
		_map2[string(i)]++
	}
	for key, val := range _map1 {
		if val == 1 && _map2[key] == 1 {
			res++
		}
	}
	return res
}