func abs(num int) int {    
    if num < 0 {
        return -1 * num
    }
    return num
}

func splitArray(nums []int) int64 {
    N := len(nums)

    prefix := make([]int, N)
    suffix := make([]int, N)
    prefbool := make([]bool, N)
    suffbool := make([]bool, N)
    
    copy(prefix, nums)
    copy(suffix, nums)

    prefbool[0] = true
    suffbool[N - 1] = true
    
    
    var prefOk, suffOk bool = true, true
    
    for i := 1; i < N; i++ {
        prefix[i] += prefix[i - 1]
        
        if nums[i - 1] >= nums[i] {
            prefOk = false
        }
        
        prefbool[i] = prefOk
        
    }
    
    for i := N - 2; i > -1; i-- {
        suffix[i] += suffix[i + 1]
        if nums[i] <= nums[i + 1] {
            suffOk = false
        }
        suffbool[i] = suffOk
    }
    
    var maxInt int64 = math.MaxInt64
    var res int64 = maxInt
    
    for i := 0; i < N - 1; i++ {
        if prefbool[i] && suffbool[i + 1] {
            res = min(res, int64(abs(prefix[i] - suffix[i + 1])))
        }
    }

    if res ^ maxInt == 0 {
        return -1
    }

    return res
}