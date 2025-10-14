func maxIncreasingSubarrays(nums []int) int {

    var (
        prev int = nums[0]
        start bool
        res, cnt, gap, oldGap, oldBlock int
    )
    res = 1
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
        res = max(res, cnt / 2)
        if oldGap == 1 {
            res = max(res, min(oldBlock, cnt))
        }
        
        prev = curr
    }
    
    return res
}