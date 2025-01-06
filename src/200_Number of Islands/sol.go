package sol

func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])

	cnt := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				cnt++
				dfs(grid, m, n, i, j)
			}
		}
	}
	return cnt
}

func dfs(grid [][]byte, m, n, idx, idy int) {
	if !inArea(idx, idy, m, n) {
		return
	}
	if grid[idx][idy] != '1' {
		return
	}
	grid[idx][idy] = '2'
	dfs(grid, m, n, idx-1, idy)
	dfs(grid, m, n, idx, idy-1)
	dfs(grid, m, n, idx+1, idy)
	dfs(grid, m, n, idx, idy+1)
}

func inArea(i, j, m, n int) bool {
	return i >= 0 && i < m && j >= 0 && j < n
}
