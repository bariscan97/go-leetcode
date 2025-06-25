func isLongPressedName(name string, typed string) bool {
	var j, cnt int
	for i := 0; i < len(name); i++ {
		if j == len(typed) {
			return false
		}
		if name[i] != typed[j] {
			i--
		} else {
			cnt++
		}
		j++
	}
	return cnt == len(name)
}