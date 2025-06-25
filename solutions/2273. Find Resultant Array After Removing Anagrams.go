func removeAnagrams(words []string) []string {
	var stack = []string{}

	for _, word := range words {
		if len(stack) == 0 || (len(stack) > 0 && sortString(word) != sortString((stack[len(stack)-1]))) {
			stack = append(stack, word)
		}
	}

	return stack
}
func sortString(s string) string {

	chars := strings.Split(s, "")

	sort.Strings(chars)

	return strings.Join(chars, "")
}