func topKFrequent(nums []int, k int) []int {
	dic := make(map[int]int)
	res := []int{}
	for _, num := range nums {
		dic[num]++
	}
	arr := [][]int{}
	for key, val := range dic {
		arr = append(arr, []int{key, val})
	}
	sort.Slice(arr, func(i int, j int) bool {
		return arr[i][1] < arr[j][1]
	})
	cnt := 0
	for i := len(arr) - 1; i >= 0; i-- {
		res = append(res, arr[i][0])
		cnt++
		if cnt == k {
			break
		}
	}
	return res
}
