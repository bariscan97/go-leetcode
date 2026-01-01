func plusOne(digits []int) []int {
	N := len(digits)
	arr := make([]int, N+1)

	digits[N-1]++

	carry := 0

	for i := N - 1; i >= 0; i-- {
		total := digits[i] + carry
		arr[i+1] = total % 10
		if total >= 10 {
			carry = 1
		} else {
			carry = 0
		}
	}

	if carry == 1 {
		arr[0] = 1
		return arr
	}

	return arr[1:]
}
