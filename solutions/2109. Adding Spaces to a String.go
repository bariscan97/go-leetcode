func addSpaces(s string, spaces []int) string {
	var (
		idx int
		res []rune
	)
	for i, chr := range s {
		if idx < len(spaces) && i == spaces[idx] {
			res = append(res, ' ')
			idx++
		}
		res = append(res, chr)
	}
	return string(res)
}