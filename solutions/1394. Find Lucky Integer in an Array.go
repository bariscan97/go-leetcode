func findLucky(arr []int) int {
    var (
        dic map[int]int = make(map[int]int)
        res int = -1
    )
    for _, i := range arr {
        dic[i]++
    }
    for i, j := range dic {
        if i == j && i > res {
            res = i
        }
    }
    return res
}