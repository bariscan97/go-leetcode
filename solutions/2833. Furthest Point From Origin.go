func furthestDistanceFromOrigin(moves string) int {
	var L, R, space int
	for _, move := range moves {
		if move == 'L' {
			L++
		} else if move == 'R' {
			R++
		} else {
			space++
		}
	}
	if L > R {
		return space + L - R
	}
	return space + R - L
}