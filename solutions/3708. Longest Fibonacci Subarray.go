func longestSubarray(nums []int) int {
    var (
        res, curr int
        cnt int = 1
        N int = len(nums)
    )
    
    for i := 1; i  < N; i++ {
        if curr == nums[i] {
            cnt++
        }else {
            cnt = 1
        }
        res = max(res, cnt)
        curr = nums[i] + nums[i - 1]        
    }
    
    return res + 1
}