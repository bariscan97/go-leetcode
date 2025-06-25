func minimumDeletions(s string) int {
	var (
		a int
		b int
	)
	for _, i := range s {
		if i == 'a' {
			a++
		} else {
			b++
		}
		if b < a {
			a = b
		}
	}
	return a
}