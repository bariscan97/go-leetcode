func minimumRecolors(blocks string, k int) int {
	dic := make(map[string]int)
	res := k
	for index, color := range blocks {
		dic[string(color)]++
		if k-1 <= index {
			diff := k - dic[string("B")]
			if diff < res {
				res = diff
			}
			dic[string(blocks[index-(k-1)])]--
		}
	}
	return res
}