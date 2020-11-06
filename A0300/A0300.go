/**
给定一个无序的整数数组，找到其中最长上升子序列的长度。

示例:

输入: [10,9,2,5,3,7,101,18]
输出: 4
解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
说明:

可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
你算法的时间复杂度应该为 O(n2) 。
进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?

通过次数164,694提交次数363,487

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-increasing-subsequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 3, 6, 7, 9, 4, 10, 5, 6}
	fmt.Println(lengthOfLIS(nums))
}

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return 1
	}

	result := make([]int, 0, len(nums))
	result = append(result, math.MaxInt32)
	R := 0
	L := 0

	for _, v := range nums {
		R = len(result)
		L = 0
		for L < R {
			mid := (L + R) / 2
			if result[mid] < v {
				L = mid + 1
			} else {
				R = mid
			}
		}
		//下面三行切片操作有问题，没办法只修改切片中的某一个元素而不影响其他元素，
		//这道题的解法看Java实现
		remainder := result[L+1:]
		result = append(result[:L], v)
		result = append(result, remainder...)
		if R == len(result) {
			result = append(result, v)
		}
	}
	return len(result)
}
