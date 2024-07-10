func minOperations(logs []string) int {
	stack := []string{}
	for _, i := range logs {
		if i == "../" {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
			continue
		}
		if i != "./" {
			stack = append(stack, i)
		}
	}
	return len(stack)
}