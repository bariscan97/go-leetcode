func countIslands(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		paths := [4][2]int{
			{i + 1, j},
			{i - 1, j},
			{i, j + 1},
			{i, j - 1},
		}
		val := grid[i][j]
		grid[i][j] = 0
		for _, path := range paths {
			r, c := path[0], path[1]
			if -1 < r && r < m && -1 < c && c < n && grid[r][c] != 0 {
				val += dfs(r, c)
			}
		}
		return val
	}
	var res int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 0 {
				total := dfs(i, j)
				if total%k == 0 {
					res++
				}
			}
		}
	}
	return res
}