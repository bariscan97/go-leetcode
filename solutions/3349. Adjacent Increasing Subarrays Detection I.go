func hasIncreasingSubarrays(nums []int, k int) bool {
    if k == 1 {
        if len(nums) > 1 {
            return true
        }
        return false
    }

    var (
        prev int = nums[0]
        start bool
        cnt, gap, oldGap, oldBlock int

    )

    for _, curr := range nums {
        if prev < curr {
            cnt++
            start = true
            if gap > 0 {
                oldGap = gap
            }
            gap = 0
        }else {
            if cnt > 1 {
                oldBlock = cnt
            }
            cnt = 1
            if start {
                gap++
            }
          
        }
        
        if cnt / 2 == k || (cnt >= k && oldBlock >= k && oldGap == 1) {
            return true
        }
        
        prev = curr
    }

    return false
}