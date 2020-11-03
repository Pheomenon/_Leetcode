package main

import (
	"fmt"
	"math"
)

/**
给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的连续子数组，并返回其长度。
如果不存在符合条件的子数组，返回 0。

示例：

输入：s = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。

进阶：

如果你已经完成了 O(n) 时间复杂度的解法, 请尝试 O(n log n) 时间复杂度的解法。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-size-subarray-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func minSubArrayLen(s int, nums []int) int {
	length := len(nums)
	start, end := 0, 0
	result := math.MaxInt32
	sum := 0
	for end < length {
		sum += nums[end]
		for sum >= s {
			sum += nums[end]
			result = min(result, end-start+1)
			start++
		}
		end++
	}

	if result == math.MaxInt32 {
		return 0
	}
	return result
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	s := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen(s, nums))
}
