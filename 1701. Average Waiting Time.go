func averageWaitingTime(customers [][]int) float64 {
	cur_time := 0
	wait_time := 0

	for _, customer := range customers {
		arv_time := customer[0]
		prep_time := customer[1]

		if cur_time <= arv_time {
			wait_time += prep_time
			cur_time = arv_time + prep_time
		} else {
			cur_time += prep_time
			wait_time += cur_time - arv_time
		}
	}

	return float64(wait_time) / float64(len(customers))
}