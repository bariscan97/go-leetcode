func maxPower(s string) int {
	val := string(s[0])
	cnt, res := 1, 0
	for _, i := range s[1:] {
		if string(i) == val {
			cnt += 1
		} else {
			cnt = 1
			val = string(i)
		}
		if cnt > res {
			res = cnt
		}

	}
	return max(res, cnt)
}