func partitionLabels(s string) []int {
	dic := make(map[rune]int)
	for _, i := range s {
		dic[i]++
	}
	var (
		cnt, size int
		res       []int
	)
	curr_dic := make(map[rune]bool)
	for _, i := range s {
		dic[i]--
		curr_dic[i] = true
		size++
		if dic[i] == 0 {
			cnt++
		}
		if cnt == len(curr_dic) {
			curr_dic = make(map[rune]bool)
			res = append(res, size)
			cnt, size = 0, 0
		}
	}
	return res
}