func numIdenticalPairs(nums []int) int {
    dic := make(map[int]int)
    var res int
    for _, num := range nums {
        dic[num]++
        res += dic[num] - 1
    }
    return res
}