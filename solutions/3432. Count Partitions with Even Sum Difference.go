func countPartitions(nums []int) int {
    N := len(nums)
    var res, sum, prefix int  

    for _, num := range nums {
        sum += num
    }   

    for _, num := range nums[:N - 1] {
        if (sum - prefix * 2) % 2 == 0 {
            res++
        }
        prefix += num
    }

    return res
}