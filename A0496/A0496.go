package main

import "math"

/*
给定两个 没有重复元素 的数组 nums1 和 nums2，其中nums1是nums2的子集。找到nums1中每个元素在nums2中的下一个比其大的值。

nums1中数字x的下一个更大元素是指x在nums2中对应位置的右边的第一个比x大的元素。如果不存在，对应位置输出 -1 。

示例 1:

输入: nums1 = [4,1,2], nums2 = [1,3,4,2].
输出: [-1,3,-1]
解释:
    对于num1中的数字4，你无法在第二个数组中找到下一个更大的数字，因此输出 -1。
    对于num1中的数字1，第二个数组中数字1右边的下一个较大数字是 3。
    对于num1中的数字2，第二个数组中没有下一个更大的数字，因此输出 -1。

示例 2:
输入: nums1 = [2,4], nums2 = [1,2,3,4].
输出: [3,-1]
解释:
    对于 num1 中的数字 2 ，第二个数组中的下一个较大数字是 3 。
	对于 num1 中的数字 4 ，第二个数组中没有下一个更大的数字，因此输出 -1 。
提示：

nums1和nums2中所有元素是唯一的。
nums1和nums2的数组大小都不超过1000。
通过次数57,852提交次数87,085

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/next-greater-element-i
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	nums1 := []int{}
	nums2 := []int{}
	nextGreaterElement(nums1, nums2)
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	size := math.MinInt64
	for i := 0; i < len(nums2); i++ {
		if size < nums2[i] {
			size = nums2[i]
		}
	}
	if size == math.MinInt64 {
		size = 0
	}
	matrix := make([][]int, size+1)
	for i := 0; i < len(nums2); i++ {
		matrix[nums2[i]] = make([]int, 0)
		for j := i + 1; j < len(nums2); j++ {
			if nums2[j] > nums2[i] {
				matrix[nums2[i]] = append(matrix[nums2[i]], nums2[j])
				break
			}
		}
	}

	ans := make([]int, len(nums1))
	for j := 0; j < len(nums1); j++ {
		if len(matrix[nums1[j]]) == 0 {
			ans[j] = -1
			continue
		}
		ans[j] = matrix[nums1[j]][0]
	}
	return ans
}
