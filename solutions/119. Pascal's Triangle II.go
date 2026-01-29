func getRow(rowIndex int) []int {
    var prev []int
    
    for i := 1; i <= rowIndex + 1; i++ {
        curr := make([]int, i)
        curr[0] = 1
        curr[len(curr) - 1] = 1
        for j := 0; j < len(prev) - 1; j++ {
            curr[j + 1] = prev[j] + prev[j + 1]
        }
        prev = curr  
    }

    return prev
}