func checkIfExist(arr []int) bool {
	_map := map[int]bool{}
	for _, v := range arr {
		if _map[2*v] || (_map[v/2] && v%2 == 0) {
			return true
		}
		_map[v] = true
	}

	return false
}