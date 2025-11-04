func findSubarrays(nums []int) bool {
    dic := make(map[int]bool)
    for i := 1; i < len(nums); i++ {
        sum := nums[i] + nums[i - 1]
        if dic[sum] {
            return true
        }
        dic[sum] = true
    }
    return false   
}