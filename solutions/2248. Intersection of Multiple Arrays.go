func intersection(nums [][]int) []int {
    dic := make(map[int]int)
    N := len(nums)
    res := []int{}
    for _, arr := range nums {
        curr := make(map[int]bool)
        for _, i := range arr {
            if !curr[i] {
                dic[i]++
                if dic[i] == N {
                    res = append(res, i)
                }
            }
            curr[i] = true
        }
    }
    sort.Ints(res)
    return res
}