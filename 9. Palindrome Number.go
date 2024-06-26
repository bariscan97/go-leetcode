func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	num := strconv.Itoa(x)
	lenNum := len(num)
	for i := 0; i < lenNum/2; i++ {
		if num[i] != num[lenNum-1-i] {
			return false
		}

	}
	return true
}