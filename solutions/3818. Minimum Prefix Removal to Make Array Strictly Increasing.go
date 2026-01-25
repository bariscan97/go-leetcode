func minimumPrefixLength(nums []int) int {
    N := len(nums)
    if N == 1 {
        return 0
    }
    for i := N - 2; i > - 1; i-- {
        if nums[i] >= nums[i + 1] {
            return i + 1
        }
    }  
    return 0
}