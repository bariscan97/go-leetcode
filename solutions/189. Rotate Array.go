func rotate(nums []int, k int)  {
    var size int = len(nums)
    var mod int = k % size
    copy(nums, append(nums[size - mod:], nums[:size - mod]...))
}