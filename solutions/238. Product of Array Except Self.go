func productExceptSelf(nums []int) []int {
    var size int = len(nums)
    arr := make([]int, size)
    prefix := []int{1,nums[0]}
    for i := 1; i < size; i++ {
        prefix = append(prefix, prefix[len(prefix) - 1] * nums[i])
    }
    var suff int = 1
    for i := size - 1; i > -1; i-- {
        arr[i] = suff * prefix[i]
        suff *= nums[i]
    }
    return arr
}