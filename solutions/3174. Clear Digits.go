func clearDigits(s string) string {
	stack := make([]string, 0)
	for _, elem := range s {
		_, err := strconv.Atoi(string(elem))
		if err != nil {
			stack = append(stack, string(elem))
		} else if len(stack) > 0 {
			stack = stack[:len(stack)-1]
		}
	}
	res := ""
	for _, i := range stack {
		res += i
	}
	return res
}