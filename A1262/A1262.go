package main

import (
	"math"
)

/**
给你一个整数数组 nums，请你找出并返回能被三整除的元素最大和。

示例 1：

输入：nums = [3,6,5,1,8]
输出：18
解释：选出数字 3, 6, 1 和 8，它们的和是 18（可被 3 整除的最大和）。
示例 2：

输入：nums = [4]
输出：0
解释：4 不能被 3 整除，所以无法选出数字，返回 0。
示例 3：

输入：nums = [1,2,3,4,4]
输出：12
解释：选出数字 1, 3, 4 以及 4，它们的和是 12（可被 3 整除的最大和）。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/greatest-sum-divisible-by-three
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	nums := []int{3, 6, 5, 1, 8}
	maxSumDivThree(nums)
}

func maxSumDivThree(nums []int) int {
	state := [3]int{0, math.MinInt64, math.MinInt64}

	for _, n := range nums {
		if n%3 == 0 {
			state[0] += n
			state[1] += n
			state[2] += n
		} else if n%3 == 2 {
			b := max(state[2]+n, state[1])
			c := max(state[0]+n, state[2])
			a := max(state[1]+n, state[0])
			state[0] = a
			state[1] = b
			state[2] = c
		} else if n%3 == 1 {
			a := max(state[2]+n, state[0])
			b := max(state[0]+n, state[1])
			c := max(state[1]+n, state[2])
			state[0] = a
			state[1] = b
			state[2] = c
		}
	}
	return state[0]
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
