func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	dic := make(map[int]int)
	for _, i := range append(items1, items2...) {
		dic[i[0]] += i[1]
	}
	res := [][]int{}
	for key, val := range dic {
		res = append(res, []int{key, val})
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i][0] < res[j][0]
	})
	return res
}