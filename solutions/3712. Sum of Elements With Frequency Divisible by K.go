func sumDivisibleByK(nums []int, k int) int {
    dic := make(map[int]int)
    var res int
    for _, num := range nums {
        if dic[num] > 0 && dic[num] % k == 0 {
            res -= dic[num] * num
        }
        dic[num]++
        if dic[num] % k == 0 {
            res += dic[num] * num
        }
    }
    return res
}