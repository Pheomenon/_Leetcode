package main

import "fmt"

/*
根据百度百科，生命游戏，简称为生命，是英国数学家约翰·何顿·康威在 1970 年发明的细胞自动机。

给定一个包含 m × n 个格子的面板，每一个格子都可以看成是一个细胞。
每个细胞都具有一个初始状态：1 即为活细胞（live），或 0 即为死细胞（dead）。
每个细胞与其八个相邻位置（水平，垂直，对角线）的细胞都遵循以下四条生存定律：

如果活细胞周围八个位置的活细胞数少于两个，则该位置活细胞死亡；
如果活细胞周围八个位置有两个或三个活细胞，则该位置活细胞仍然存活；
如果活细胞周围八个位置有超过三个活细胞，则该位置活细胞死亡；
如果死细胞周围正好有三个活细胞，则该位置死细胞复活；

下一个状态是通过将上述规则同时应用于当前状态下的每个细胞所形成的，
其中细胞的出生和死亡是同时发生的。给你 m x n 网格面板 board 的当前状态，返回下一个状态。

示例 1：
输入：board = [[0,1,0],[0,0,1],[1,1,1],[0,0,0]]
输出：[[0,0,0],[1,0,1],[0,1,1],[0,1,0]]

示例 2：
输入：board = [[1,1],[1,0]]
输出：[[1,1],[1,1]]

提示：

m == board.length
n == board[i].length
1 <= m, n <= 25
board[i][j] 为 0 或 1

进阶：

你可以使用原地算法解决本题吗？请注意，面板上所有格子需要同时被更新：你不能先更新某些格子，然后使用它们的更新后的值再更新其他格子。
本题中，我们使用二维数组来表示面板。原则上，面板是无限的，但当活细胞侵占了面板边界时会造成问题。你将如何解决这些问题？
通过次数42,928提交次数57,399

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/game-of-life
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	matrix := [][]int{
		{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0},
	}
	gameOfLife(matrix)
	for i := 0; i < len(matrix); i++ {
		for k := 0; k < len(matrix[i]); k++ {
			fmt.Print(matrix[i][k])
		}
	}
}

func gameOfLife(board [][]int) {
	ans := make([][]int, len(board))

	for i := 0; i < len(board); i++ {
		tmp := make([]int, len(board[i]))
		ans[i] = append(ans[i], tmp...)
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			ans[i][j] = board[i][j]
		}
	}

	near := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			board[i][j] = 0
		}
	}

	for i := 0; i < len(ans); i++ {
		for k := 0; k < len(ans[i]); k++ {
			near = dfs(i+1, k, ans) + dfs(i-1, k, ans) + dfs(i, k+1, ans) + dfs(i, k-1, ans) + dfs(i+1, k+1, ans) + dfs(i+1, k-1, ans) + dfs(i-1, k+1, ans) + dfs(i-1, k-1, ans)
			if near < 2 {
				board[i][k] = 0
			} else if near == 2 && ans[i][k] == 1 || near == 3 && ans[i][k] == 1 {
				board[i][k] = 1
			} else if near > 3 {
				board[i][k] = 0
			} else if board[i][k] == 0 && near == 3 {
				board[i][k] = 1
			}
		}
	}
}

func dfs(x, y int, board [][]int) int {
	near := 0
	if x >= 0 && x < len(board) && y >= 0 && y < len(board[x]) {
		if board[x][y] == 1 {
			near++
			return near
		} else {
			return near
		}
	}
	return near
}
