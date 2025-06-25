func topKFrequent(words []string, k int) []string {
	_map := make(map[string]int)
	_map2 := make(map[int][]string)
	res := make([]string, 0)
	for _, word := range words {
		_map[word]++
	}
	for key, val := range _map {
		_map2[val] = append(_map2[val], key)
	}
	vals := make([]int, 0)
	for key, val := range _map2 {
		vals = append(vals, key)
		sort.Strings(val)
	}
	sort.Ints(vals)
	cnt := 0
	for i := len(vals) - 1; i >= 0; i-- {
		flag := true
		for _, word := range _map2[vals[i]] {
			res = append(res, word)
			cnt++
			if cnt == k {
				flag = false
				break
			}
		}
		if !flag {
			break
		}
	}
	return res
}