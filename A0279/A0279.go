package main

import "fmt"

/*
给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。

示例 1:
输入: n = 12
输出: 3
解释: 12 = 4 + 4 + 4.

示例 2:
输入: n = 13
输出: 2
解释: 13 = 4 + 9.
通过次数103,610提交次数176,248

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/perfect-squares
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	n := 13
	fmt.Println(numSquares(n))
}

func numSquares(n int) int {
	perfectSquare := make([]int, 0)
	now := 1
	for i := 0; now <= n; i++ {
		now = pow(i + 1)
		perfectSquare = append(perfectSquare, now)
	}

	dp := make([]int, n+1)

	for i := 1; i < len(dp); i++ {
		dp[i] = n + 1
	}

	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(perfectSquare)-1; j++ {
			if i-perfectSquare[j] < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[i-perfectSquare[j]]+1)
		}
	}
	return dp[len(dp)-1]
}

func pow(x int) int {
	return x * x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
