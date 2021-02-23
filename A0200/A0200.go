package main

import "fmt"

/**
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
此外，你可以假设该网格的四条边均被水包围。

示例 1：
输入：grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
输出：1

示例 2：
输入：grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
输出：3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-islands
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	item := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	fmt.Println(numIslands(item))
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	var ans int
	var dfs func(r, l int)
	dfs = func(r, l int) {
		if !inArea(grid, r, l) {
			return
		}
		if grid[r][l] == '1' {
			grid[r][l] = 0
			dfs(r+1, l)
			dfs(r-1, l)
			dfs(r, l+1)
			dfs(r, l-1)
		} else {
			return
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				ans++
				dfs(i, j)
			}
		}
	}
	return ans
}

func inArea(grid [][]byte, row, col int) bool {
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) {
		return false
	}
	return true
}
