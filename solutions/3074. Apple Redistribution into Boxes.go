func minimumBoxes(apple []int, capacity []int) int {
	var total int
	for _, i := range apple {
		total += i
	}
	sort.Ints(capacity)
	for i := len(capacity) - 1; i > -1; i-- {
		total -= capacity[i]
		if total <= 0 {
			return len(capacity) - i
		}
	}
	return len(capacity)
}