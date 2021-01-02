package main

/*
给你一个长度为n的整数数组nums，其中n > 1，返回输出数组output，其中 output[i]等于nums中除nums[i]之外其余各元素的乘积。

示例:

输入: [1,2,3,4]
输出: [24,12,8,6]

提示：题目数据保证数组之中任意元素的全部前缀元素和后缀（甚至是整个数组）的乘积都在 32 位整数范围内。

说明: 请不要使用除法，且在O(n) 时间复杂度内完成此题。

进阶：
你可以在常数空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组不被视为额外空间。）

通过次数91,406提交次数128,306

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/product-of-array-except-self
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {

}

//创建一个左乘积数组和一个右乘积数组
//左右的起点分别是1
//ans的结果就是left[i]+right[i]
func productExceptSelf(nums []int) []int {
	right, left, ans := make([]int, len(nums)), make([]int, len(nums)), make([]int, len(nums))
	left[0], right[len(right)-1] = 1, 1
	for i := 1; i < len(left); i++ {
		left[i] = left[i-1] * nums[i-1]
	}
	for i := len(right) - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}
	for i := 0; i < len(nums); i++ {
		ans[i] = left[i] * right[i]
	}
	return ans
}
