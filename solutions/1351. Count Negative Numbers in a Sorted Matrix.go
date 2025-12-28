func countNegatives(grid [][]int) int {
    var res int

    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] < 0 {
                res++
            }
        }
    }

    return res
}