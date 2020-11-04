package main

import (
	"fmt"
)

/**
在二维地图上， 0代表海洋， 1代表陆地，我们最多只能将一格 0 海洋变成 1变成陆地。
进行填海之后，地图上最大的岛屿面积是多少？（上、下、左、右四个方向相连的 1 可形成岛屿）
示例 1:
输入: [[1, 0], [0, 1]]
输出: 3
解释: 将一格0变成1，最终连通两个小岛得到面积为 3 的岛屿。
示例 2:

输入: [[1, 1], [1, 0]]
输出: 4
解释: 将一格0变成1，岛屿的面积扩大为 4。
示例 3:

输入: [[1, 1], [1, 1]]
输出: 4
解释: 没有0可以让我们变成1，面积依然为 4。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/making-a-large-island
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	tmp := [][]int{{0, 0, 0, 0, 0, 0, 0},
		{0, 1, 1, 1, 1, 0, 0},
		{0, 1, 0, 0, 1, 0, 0},
		{1, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 0, 0},
		{0, 1, 1, 1, 1, 0, 0}}
	fmt.Println(largestIsland(tmp))
}

//用于存储每个岛的大小
var landSize [256]int

func largestIsland(grid [][]int) int {
	r := len(grid)
	c := len(grid[0])

	check := 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			check = max(grid[i][j], check)
		}
	}
	if check == 0 {
		return 1
	}

	if check == 1 && r == 1 {
		return 1
	}
	if check == 1 && c == 1 {
		return 1
	}

	tag := 2
	result := 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if grid[i][j] == 1 {
				landSize[tag] = landDFS(grid, i, j, tag)
				tag++
			}
		}
	}

	count := 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if grid[i][j] != 0 {
				count++
			}
		}
	}

	if count == r*c {
		return r * c
	}
	if count == r*c-1 {
		return r * c
	}
	if count == 0 {
		return 1
	}
	if count == 1 && r*c == 1 {
		return 1
	}
	if count == 1 && r*c != 1 {
		return 2
	}

	//遍历海洋格子
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if grid[i][j] == 0 {
				result = max(detect(grid, i, j)+1, result)
			}
		}
	}
	return result
}

/**
深度遍历图，把每个岛屿打上tag，把原本是1的位置为tag，这个tag就是LandSize的索引
*/
func landDFS(grid [][]int, r, c, tag int) int {
	if !inArea(grid, r, c) {
		return 0
	}
	if grid[r][c] == 0 {
		return 0
	}
	if grid[r][c] != 1 {
		return 0
	}
	grid[r][c] = tag
	return 1 + landDFS(grid, r-1, c, tag) + landDFS(grid, r+1, c, tag) + landDFS(grid, r, c-1, tag) + landDFS(grid, r, c+1, tag)
}

//累加海洋格子周围的陆地格子的面积
func detect(grid [][]int, r, c int) int {
	tags := make([]int, 4)
	tag := 0
	result := 0
	if inArea(grid, r+1, c) {
		if grid[r+1][c] != 0 {
			tag = grid[r+1][c]
			tags = append(tags, tag)
			result += landSize[tag]
		}
	}
	if inArea(grid, r-1, c) {
		tag = grid[r-1][c]
		if !exist(tags, tag) {
			if grid[r-1][c] != 0 {
				tags = append(tags, tag)
				result += landSize[tag]
			}
		}
	}
	if inArea(grid, r, c+1) {
		tag = grid[r][c+1]
		if !exist(tags, tag) {
			if grid[r][c+1] != 0 {
				tags = append(tags, tag)
				result += landSize[tag]
			}
		}
	}
	if inArea(grid, r, c-1) {
		tag = grid[r][c-1]
		if !exist(tags, tag) {
			if grid[r][c-1] != 0 {
				tags = append(tags, tag)
				result += landSize[tag]
			}
		}
	}
	return result
}

func inArea(grid [][]int, r, c int) bool {
	return 0 <= r && r < len(grid) && 0 <= c && c < len(grid[0])
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func exist(nums []int, num int) bool {
	for i := 0; i < len(nums); i++ {
		if num == nums[i] {
			return true
		}
	}
	return false
}
