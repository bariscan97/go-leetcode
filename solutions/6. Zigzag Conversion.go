func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	var b strings.Builder
	b.Grow(len(s))
	step := numRows*2 - 2

	for row := 0; row < numRows; row++ {
		nextStep := 0
		if row == 0 || row == numRows-1 {
			nextStep = step
		} else {
			nextStep = row * 2
		}
		for pos := row; pos < len(s); pos += nextStep {
			b.WriteByte(s[pos])
			if row == 0 || row == numRows-1 {
				nextStep = step
			} else {
				nextStep = step - nextStep
			}
		}
	}

	return b.String()
}