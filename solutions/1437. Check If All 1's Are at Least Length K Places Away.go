func kLengthApart(nums []int, k int) bool {
    zeroCnt := 1_000_00
    for _, num := range nums {
        if num == 0 {
            zeroCnt++
        }else {
            if zeroCnt < k {
                
                return false
            }
            zeroCnt = 0
        }
    }
    return true
}