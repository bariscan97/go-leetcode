func countHillValley(nums []int) int {
    var (
        res int
        prev int = nums[0]
        flag rune = ' '
    )
    for _, num := range nums[1:] {
        if prev < num {
            if flag == 'v' {
                res++
            }
            flag = 'h'
            
        }else if prev > num {
            if flag == 'h' {
                res++
            }
            flag = 'v'
            
        }
        prev = num
    }
    
    return res
}