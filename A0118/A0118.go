package main

/*
给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
在杨辉三角中，每个数是它左上方和右上方的数的和。

示例:
输入: 5
输出:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]
通过次数146,680提交次数210,175

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/pascals-triangle
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	n := 5
	generate(n)
}

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	ans := make([][]int, numRows)
	ans[0] = append(ans[0], []int{1}...)
	if numRows == 1 {
		return ans
	}
	ans[1] = append(ans[1], []int{1, 1}...)
	if numRows == 2 {
		return ans
	}

	for i := 2; i < numRows; i++ {
		ans[i] = rowGen(ans[i-1])
	}
	return ans
}

func rowGen(row []int) []int {
	res := make([]int, len(row)+1)
	res[0] = 1
	idx := 1
	left, right := 0, 1
	for right != len(row) {
		res[idx] = row[right] + row[left]
		right++
		left++
		idx++
	}
	res[len(res)-1] = 1
	return res
}
