func numberOfPairs(nums []int) []int {
    dic := make(map[int]int)
    var cnt, odd int
    res := make([]int, 2)
    for _, num := range nums {
        dic[num]++
    }
    for _, v := range dic {
        if v % 2 == 1 {
            odd++
        }
        if v > 1 {
            cnt += v / 2
        }
    }
    res[0] = cnt
    res[1] = odd
    return res
}