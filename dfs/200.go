package dfs

/*
个人思路：和130类似，找到陆地后，对陆地进行dfs，记录结果后再转换成水
答案：可以再优化一下，用一个数组记录是否访问过，感觉优化的也不多
*/

func numIslands(grid [][]byte) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				res++
				numIslandsFunc(grid, i, j)
			}
		}
	}
	return res
}

func numIslandsFunc(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	numIslandsFunc(grid, i, j+1)
	numIslandsFunc(grid, i, j-1)
	numIslandsFunc(grid, i+1, j)
	numIslandsFunc(grid, i-1, j)
}
