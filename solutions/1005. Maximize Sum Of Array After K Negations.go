func largestSumAfterKNegations(nums []int, k int) int {
    for k > 0 {
        idx := slices.Index(nums, slices.Min(nums))
        nums[idx] = -nums[idx]
        k--
    }
    sum := 0
    for _, n := range nums {
        sum += n
    }
    return sum
}