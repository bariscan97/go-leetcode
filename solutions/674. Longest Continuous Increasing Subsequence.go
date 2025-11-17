func findLengthOfLCIS(nums []int) int {
    cnt := 1
    res := 1
    for i := 1; i < len(nums); i++ {
        if nums[i] > nums[i - 1] {
            cnt++
        }else {
            cnt = 1
        }
        res = max(res, cnt)
    }
    return res
}