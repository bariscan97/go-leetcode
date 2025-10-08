func alternatingSum(nums []int) int {
    var res int
    for i, num := range nums {
        if i % 2 == 1 {
            res -= num
        }else {
            res += num
        }
    }
    return res
}