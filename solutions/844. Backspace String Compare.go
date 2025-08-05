func isSame(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func backspace(s string) []rune {
	var stack []rune
	for _, i := range s {
		if i == '#' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, i)
		}
	}
	return stack
}

func backspaceCompare(s string, t string) bool {
	return isSame(backspace(s), backspace(t))
}