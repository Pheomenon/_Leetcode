package main

import "fmt"

/*
给定一个二维网格和一个单词，找出该单词是否存在于网格中。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

示例:

board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

给定 word = "ABCCED", 返回 true
给定 word = "SEE", 返回 true
给定 word = "ABCB", 返回 false

提示：

board 和 word 中只包含大写和小写英文字母。
1 <= board.length <= 200
1 <= board[i].length <= 200
1 <= word.length <= 10^3
通过次数128,354提交次数292,733

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/word-search
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'E', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCEFSADEESE"
	fmt.Println(exist(board, word))
}

func dfs(board [][]byte, y, x int, word string, visited map[string]int) bool {
	if y < 0 || y > len(board)-1 || x < 0 || x > len(board[0])-1 {
		return false
	}
	if word[0] == board[y][x] {
		key := fmt.Sprintf("%d_%d", y, x)
		if _, ok := visited[key]; ok {
			return false
		}
		visited[key]++
		if len(word) == 1 {
			return true
		}
		return dfs(board, y, x+1, word[1:], visited) ||
			dfs(board, y+1, x, word[1:], visited) ||
			dfs(board, y, x-1, word[1:], visited) ||
			dfs(board, y-1, x, word[1:], visited)
	} else {
		return false
	}
}

func exist(board [][]byte, word string) bool {
	if word == "" || board == nil {
		return false
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if word[0] == board[i][j] {
				visited := map[string]int{}
				if dfs(board, i, j, word, visited) {
					return true
				}
			}
		}
	}
	return false
}
