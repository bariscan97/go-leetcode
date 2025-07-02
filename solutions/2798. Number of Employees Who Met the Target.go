func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
	var res int
	for _, hour := range hours {
		if hour >= target {
			res++
		}
	}
	return res
}