func numberOfBeams(bank []string) int {
	var res, prev, curr int
	for idx, i := range bank {
		curr = 0
		for _, j := range i {
			if j == '1' {
				curr++
			}
		}
		if curr == 0 && idx != 0 {
			continue
		}
		res += prev * curr
		prev = curr
	}
	return res
}