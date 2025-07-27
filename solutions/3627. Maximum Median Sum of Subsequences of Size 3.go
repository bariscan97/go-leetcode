func maximumMedianSum(nums []int) int64 {
    sort.Ints(nums)
    var res int
    for 
    i, size := len(nums) - 2, len(nums) / 3; 
    i > -1 && size > 0 ;
    i, size = i - 2, size - 1 {
        res += nums[i]
    }
    return int64(res)
}