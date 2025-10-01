func reversed(nums []int) {
    N := len(nums)
    for i := 0; i < N / 2; i++ {
        nums[i], nums[(N - 1) - i] = nums[(N - 1) - i], nums[i]
    }
}

func decimalRepresentation(n int) []int {
    var res []int
    var digit int = 1
    
    for n > 0 {
        num := n % 10
        n /= 10
        if num != 0 {
            res = append(res, num * digit)
        }
        digit *= 10
    }
    
    reversed(res)
    
    return res
}