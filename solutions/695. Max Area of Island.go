func maxAreaOfIsland(grid [][]int) int {
	var (
		dfs func(i, j int) int
		res int
	)
	m, n := len(grid), len(grid[0])
	dfs = func(i, j int) int {
		paths := [][]int{
			{i + 1, j},
			{i - 1, j},
			{i, j + 1},
			{i, j - 1},
		}
		val := 1
		grid[i][j] = 0
		for i := range paths {
			r := paths[i][0]
			c := paths[i][1]
			if -1 < r && r < m && -1 < c && c < n && grid[r][c] == 1 {
				val += dfs(r, c)
			}
		}
		return val
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				res = max(res, dfs(i, j))
			}
		}
	}
	return res
}