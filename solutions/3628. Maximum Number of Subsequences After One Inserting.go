func max(args ...int) int {
	var res int = args[0]
	for _, arg := range args[1:] {
		if arg > res {
			res = arg
		}
	}
	return res
}
func numOfSubsequences(s string) int64 {
	var size int = len(s)
	var a, b int
	suffix := make([]int, size)
	prefix := make([]int, size)
	for i := 0; i < size; i++ {
		if s[i] == 'L' {
			a++
		}
		if s[size-1-i] == 'T' {
			b++
		}
		prefix[i] = a
		suffix[size-1-i] = b
	}
	var x, y, z, zz int
	for i := 0; i < size; i++ {
		if rune(s[i]) == 'C' {
			x += ((prefix[i] + 1) * suffix[i])
			y += (prefix[i] * (suffix[i] + 1))
			zz += prefix[i] * suffix[i]
		}
		z = max(z, prefix[i]*suffix[i])
	}

	return int64(max(x, y, z+zz))
}