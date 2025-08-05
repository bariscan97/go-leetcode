func removeStars(s string) string {
	var stack []rune
	for _, i := range s {
		if i == '*' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, i)
		}
	}
	return string(stack)
}