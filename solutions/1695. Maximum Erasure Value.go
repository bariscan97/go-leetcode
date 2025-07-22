func maximumUniqueSubarray(nums []int) int {
    dic := make(map[int]int)
    var res, total, j int
    for _, num := range nums {
        dic[num]++
        total += num
        for dic[num] == 2 {
            dic[nums[j]]--
            total -= nums[j]
            j++
        }
        if total > res {
            res = total
        }
    }
    return res
}