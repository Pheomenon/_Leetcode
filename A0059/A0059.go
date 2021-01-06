package main

/*
给定一个正整数 n，生成一个包含 1 到 n^2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。

示例:

输入: 3
输出:
[
 [ 1, 2, 3 ],
 [ 8, 9, 4 ],
 [ 7, 6, 5 ]
]
通过次数56,208提交次数71,590

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/spiral-matrix-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	generateMatrix(4)
}

func generateMatrix(n int) [][]int {
	context := 1

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	left, right, top, bottom := 0, n-1, 0, n-1
	direct := 0

	for top <= bottom && left <= right {
		switch direct {
		case 0: // →
			for i := left; i <= right; i++ {
				matrix[top][i] = context
				context++
			}
			top++
			direct = 1
		case 1: // ↓
			for i := top; i <= bottom; i++ {
				matrix[i][right] = context
				context++
			}
			right--
			direct = 2
		case 2: // ←
			for i := right; i >= left; i-- {
				matrix[bottom][i] = context
				context++
			}
			bottom--
			direct = 3
		case 3: // ↑
			for i := bottom; i >= top; i-- {
				matrix[i][left] = context
				context++
			}
			left++
			direct = 0
		}
	}
	return matrix
}
