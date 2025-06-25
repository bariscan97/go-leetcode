func maxNumberOfBalloons(text string) int {
	_map := make(map[string]int)
	for _, v := range text {
		if _, ok := m[string(v)]; !ok {
			_map[string(v)] = 1
		} else {
			_map[string(v)] = _map[string(v)] + 1
		}
	}

	return int(math.Min(float64(_map["b"]), math.Min(float64(_map["a"]), math.Min(float64(_map["l"]/2), math.Min(float64(_map["o"]/2), float64(_map["n"]))))))
}