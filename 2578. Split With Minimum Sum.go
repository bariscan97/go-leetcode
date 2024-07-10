func splitNum(num int) int {
	num1, num2 := "", ""
	for idx, val := range sortString(strconv.Itoa(num)) {
		if idx%2 == 0 {
			num1 += string(val)
		} else {
			num2 += string(val)
		}
	}
	n1, _ := strconv.Atoi(num1)
	n2, _ := strconv.Atoi(num2)
	return n1 + n2
}

func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}