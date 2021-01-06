package main

/*
给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。
示例 1:
输入:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
输出: [1,2,3,6,9,8,7,4,5]

示例 2:
输入:
[
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9,10,11,12]
]
输出: [1,2,3,4,8,12,11,10,9,5,6,7]
通过次数96,643提交次数229,643

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/spiral-matrix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	spiralOrder(matrix)
}

func spiralOrder(matrix [][]int) []int {
	ans := make([]int, 0)
	left, right, top, bottom := 0, len(matrix[0])-1, 0, len(matrix)-1
	direct := 0

	for top <= bottom && left <= right {
		switch direct {
		case 0: // →
			for i := left; i <= right; i++ {
				ans = append(ans, matrix[top][i])
			}
			top++
			direct = 1
		case 1: // ↓
			for i := top; i <= bottom; i++ {
				ans = append(ans, matrix[i][right])
			}
			right--
			direct = 2
		case 2: // ←
			for i := right; i >= left; i-- {
				ans = append(ans, matrix[bottom][i])
			}
			bottom--
			direct = 3
		case 3: // ↑
			for i := bottom; i >= top; i-- {
				ans = append(ans, matrix[i][left])
			}
			left++
			direct = 0
		}
	}
	return ans
}
