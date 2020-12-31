package main

import (
	"fmt"
)

/*
数字n代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且有效的 括号组合。

示例：

输入：n = 3
输出：[
       "((()))",
       "(()())",
       "(())()",
       "()(())",
       "()()()"
     ]
通过次数212,196提交次数276,942

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/generate-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	n := 3
	fmt.Println(generateParenthesis(n))
}

func generateParenthesis(n int) []string {
	var track string
	brackets := []string{"(", ")"}
	ans := make([]string, 0)
	var backtrack func(track string, open, close int)
	backtrack = func(track string, open, close int) {

		if len(track) == 2*n {
			ans = append(ans, track)
			return
		}

		if open < n {
			track += brackets[0]
			backtrack(track, open+1, close)
			track = track[:len(track)-1]
		}

		if close < open {
			track += brackets[1]
			backtrack(track, open, close+1)
			track = track[:len(track)-1]
		}
	}
	backtrack(track, 0, 0)
	return ans
}
