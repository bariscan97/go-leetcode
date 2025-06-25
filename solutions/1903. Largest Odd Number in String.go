func largestOddNumber(num string) string {
	for i := len(num) - 1; i >= 0; i-- {
		val, _ := strconv.Atoi(string(num[i]))
		if val%2 == 1 {
			return num[:i+1]
		}
	}
	return ""
}