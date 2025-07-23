func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func calc(stack []rune, x, y int) int {
	if len(stack) == 0 {
		return 0
	}

	val := stack[0]
	cnt := 1
	n := len(stack)

	for _, r := range stack[1:] {
		if r == val {
			cnt++
		} else {
			diff := n - cnt
			if val == 'b' {
				return min(cnt, diff) * y
			}
			return min(cnt, diff) * x
		}
	}
	return 0
}

func maximumGain(s string, x, y int) int {
	var stack []rune
	res := 0

	for _, r := range s {
		if r != 'a' && r != 'b' {
			res += calc(stack, x, y)
			stack = stack[:0]
			continue
		}

		if x > y {
			if len(stack) > 0 && stack[len(stack)-1] == 'a' && r == 'b' {
				stack = stack[:len(stack)-1]
				res += x
				continue
			}
		} else {
			if len(stack) > 0 && stack[len(stack)-1] == 'b' && r == 'a' {
				stack = stack[:len(stack)-1]
				res += y
				continue
			}
		}
		stack = append(stack, r)
	}
	res += calc(stack, x, y)
	return res
}