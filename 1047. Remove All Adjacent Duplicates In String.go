func removeDuplicates(s string) string {
	stack := make([]rune, 0)

	for _, elem := range s {
		if len(stack) != 0 && stack[len(stack)-1] == elem {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, elem)
		}
	}

	return string(stack)
}