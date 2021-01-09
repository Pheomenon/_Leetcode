package main

/*
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。

求在该柱状图中，能够勾勒出来的矩形的最大面积。

以上是柱状图的示例，其中每个柱子的宽度为 1，给定的高度为 [2,1,5,6,2,3]。
图中阴影部分为所能勾勒出的最大矩形面积，其面积为 10 个单位。
示例:

输入: [2,1,5,6,2,3]
输出: 10
通过次数113,495提交次数267,198

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/largest-rectangle-in-histogram
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	height := []int{4, 2}
	largestRectangleArea(height)
}

//单调栈

func largestRectangleArea(heights []int) int {
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)
	stack := make([]int, 0)
	stack = append(stack, 0)
	ptr := 0
	ans := 0

	for i := 1; i < len(heights); i++ {
		for ptr != 0 && heights[i] < heights[stack[ptr]] {
			h := heights[stack[ptr]]
			ptr--
			ans = max(ans, (i-stack[ptr]-1)*h)
			stack = stack[:len(stack)-1]

		}
		ptr++
		stack = append(stack, i)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
