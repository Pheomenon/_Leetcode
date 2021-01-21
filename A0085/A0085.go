package main

import "fmt"

/*
给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。

示例 1：
输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
输出：6
解释：最大矩形如上图所示。

示例 2：
输入：matrix = []
输出：0

示例 3：
输入：matrix = [["0"]]
输出：0

示例 4：
输入：matrix = [["1"]]
输出：1

示例 5：
输入：matrix = [["0","0"]]
输出：0

提示：

rows == matrix.length
cols == matrix[0].length
0 <= row, cols <= 200
matrix[i][j] 为 '0' 或 '1'
通过次数65,345提交次数126,604

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximal-rectangle
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
		//{'0', '1'}, {'0', '1'},
		//{'0', '1'}, {'1', '0'},
	}
	fmt.Println(maximalRectangle(matrix))
}

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	para := make([]int, len(matrix[0]))
	ans := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if i-1 >= 0 && matrix[i][j] != '0' {
				matrix[i][j] += matrix[i-1][j]
				matrix[i][j] -= '0'
				para[j] = int(matrix[i][j])
			} else {
				matrix[i][j] = matrix[i][j] - '0'
				para[j] = int(matrix[i][j])
			}
		}
		ans = max(ans, check(para))
	}
	return ans
}

func check(row []int) int {
	res := 0
	stack := make([]int, 0)
	ptr := 0
	row = append([]int{0}, row...)
	row = append(row, 0)
	stack = append(stack, 0)

	for i := 1; i < len(row); i++ {
		for ptr != 0 && row[i] < row[stack[ptr]] {
			h := row[stack[ptr]]
			ptr--
			res = max(res, (i-stack[ptr]-1)*h)
			stack = stack[:len(stack)-1]
		}
		ptr++
		stack = append(stack, i)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
