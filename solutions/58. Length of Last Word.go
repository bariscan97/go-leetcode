func lengthOfLastWord(s string) int {
	flag := false
	res := 0
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == " " && !flag {
			continue
		}
		if string(s[i]) != " " {
			flag = true
			res++
		}
		if string(s[i]) == " " {
			break
		}
	}
	return res
}