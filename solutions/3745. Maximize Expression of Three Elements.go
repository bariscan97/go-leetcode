func maximizeExpressionOfThree(nums []int) int {
    sort.Ints(nums)
    N := len(nums)
    return nums[N - 1] + nums[N - 2] - nums[0]
}