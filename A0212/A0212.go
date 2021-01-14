package main

/*
给定一个 m x n 二维字符网格board和一个单词（字符串）列表 words，
找出所有同时在二维网格和字典中出现的单词。
单词必须按照字母顺序，通过相邻的单元格内的字母构成，
其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。
同一个单元格内的字母在一个单词中不允许被重复使用。

示例 1：
输入：board = [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]], words = ["oath","pea","eat","rain"]
输出：["eat","oath"]

示例 2：
输入：board = [["a","b"],["c","d"]], words = ["abcb"]
输出：[]

提示：

m == board.length
n == board[i].length
1 <= m, n <= 12
board[i][j] 是一个小写英文字母
1 <= words.length <= 3 * 104
1 <= words[i].length <= 10
words[i] 由小写英文字母组成
words 中的所有字符串互不相同
通过次数27,517提交次数61,698

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/word-search-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'}}
	words := []string{"oath", "pea", "eat", "rain"}
	findWords(board, words)

}

func findWords(board [][]byte, words []string) []string {
	ans := make([]string, 0)
	maxlength := 0
	for _, word := range words {
		if maxlength < len(word) {
			maxlength = len(word)
		}
	}
	var dfs func(board [][]byte, track string, row, col int, maxlength int)
	dfs = func(board [][]byte, track string, row, col int, maxlength int) {
		if row == -1 || col == -1 || row > len(board)-1 || col > len(board[0])-1 {
			return
		}
		for _, word := range words {
			if word == track {
				ans = append(ans, track)
				return
			}
		}
		if len(track) == maxlength {
			return
		}

		for i := row; i < len(board); i++ {
			for j := col; j < len(board[i]); j++ {
				track += string(board[i][j])
				dfs(board, track, row+1, col, maxlength)
				track = track[:len(track)-1]
				dfs(board, track, row, col+1, maxlength)
				track = track[:len(track)-1]
				dfs(board, track, row-1, col, maxlength)
				track = track[:len(track)-1]
				dfs(board, track, row, col-1, maxlength)
				track = track[:len(track)-1]
			}
		}
	}
	dfs(board, "", 0, 0, maxlength)
	return ans
}
