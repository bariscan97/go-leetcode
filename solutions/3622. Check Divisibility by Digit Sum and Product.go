func checkDivisibility(n int) bool {
	a, b, c := 1, 0, n
	for c > 0 {
		x := c % 10
		a *= x
		b += x
		c /= 10
	}
	return n%(a+b) == 0
}