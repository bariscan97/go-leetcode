func minimumDistance(nums []int) int {
	dic := make(map[int][]int)
	res := int(^uint(0) >> 1)
	for i, num := range nums {
		dic[num] = append(dic[num], i)
		if len(dic[num]) == 3 {
			res = min(res, (dic[num][2]-dic[num][0])*2)
			dic[num] = dic[num][1:]
		}
	}
	if res^int(^uint(0)>>1) == 0 {
		return -1
	}
	return res
}