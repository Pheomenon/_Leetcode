package main

import (
	"fmt"
)

/*
给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
说明：每次只能向下或者向右移动一步。

示例 1：
输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
输出：7
解释：因为路径 1→3→1→1→1 的总和最小。

示例 2：
输入：grid = [[1,2,3],[4,5,6]]
输出：12

提示：

m == grid.length
n == grid[i].length
1 <= m, n <= 200
0 <= grid[i][j] <= 100
通过次数173,577提交次数255,205

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-path-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	grid := [][]int{
		{1, 2, 3}, {4, 5, 6},
	}
	fmt.Println(minPathSum(grid))
}

func minPathSum(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])
	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < col; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}
	for j := 1; j < row; j++ {
		dp[j][0] = dp[j-1][0] + grid[j][0]
	}

	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			minV := min(dp[i-1][j], dp[i][j-1])
			dp[i][j] = minV + grid[i][j]
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
