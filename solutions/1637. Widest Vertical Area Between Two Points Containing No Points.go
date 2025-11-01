func maxWidthOfVerticalArea(points [][]int) int {
	maxVal := 0
	var tmp []int

	for _, v := range points {
		tmp = append(tmp, v[0])
	}

	sort.Ints(tmp)

	for i := 1; i < len(tmp); i++ {
		if tmp[i]-tmp[i-1] > maxVal {
			maxVal = tmp[i] - tmp[i-1]
		}
	}
	return maxVal
}