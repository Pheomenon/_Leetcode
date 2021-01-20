package main

import "fmt"

/*
在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。

示例 1：
输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
输出：4

示例 2：
输入：matrix = [["0","1"],["1","0"]]
输出：1

示例 3：
输入：matrix = [["0"]]
输出：0

提示：

m == matrix.length
n == matrix[i].length
1 <= m, n <= 300
matrix[i][j] 为 '0' 或 '1'
通过次数86,222提交次数196,717

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximal-square
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	matrix := [][]byte{
		//{'1', '0', '1', '0', '0'},
		//{'1', '0', '1', '1', '1'},
		//{'1', '1', '1', '1', '1'},
		//{'1', '0', '0', '1', '0'},
		//{'0', '1'},
		//{'1', '0'},
		{'0'},
	}
	fmt.Println(maximalSquare(matrix))
}

func maximalSquare(matrix [][]byte) int {
	var ans byte
	ans = '0'
	var dp func(x, y int)
	dp = func(x, y int) {
		//if x > len(matrix)-1 || x < 0 || y > len(matrix[0])-1 || y < 0 {
		//	return
		//}

		if x == 0 || y == 0 || x == len(matrix) || y == len(matrix[x]) {
			if matrix[x][y] == '1' && ans == '0' {
				ans = '1'
			}
			return
		}

		if matrix[x][y] == '1' {
			matrix[x][y] = min(matrix[x-1][y], min(matrix[x][y-1], matrix[x-1][y-1])) + 1
			ans = max(ans, matrix[x][y])
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			dp(i, j)
		}
	}
	if ans != '0' {
		ans = ans - '0'
		return int(ans) * int(ans)
	} else {
		return 0
	}
}

func max(x, y byte) byte {
	if x > y {
		return x
	}
	return y
}

func min(x, y byte) byte {
	if x < y {
		return x
	}
	return y
}
