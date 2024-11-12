func findPermutationDifference(s string, t string) int {
	var res float64
	for i, char := range s {
		res += math.Abs(float64(i - strings.Index(t, string(char))))
	}
	return int(res)
}
