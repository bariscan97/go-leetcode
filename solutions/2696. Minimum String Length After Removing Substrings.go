func minLength(s string) int {
	stack := []string{}
	for _, char := range s {
		flag := true
		if len(stack) > 0 {
			if stack[len(stack)-1] == "A" && string(char) == "B" {
				stack = stack[:len(stack)-1]
				flag = false
			} else if stack[len(stack)-1] == "C" && string(char) == "D" {
				stack = stack[:len(stack)-1]
				flag = false
			}
		}
		if flag {
			stack = append(stack, string(char))
		}

	}
	return len(stack)
}