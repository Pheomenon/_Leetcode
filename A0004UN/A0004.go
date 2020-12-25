package main

/*
给定两个大小为 m 和 n 的正序（从小到大）数组nums1 和nums2。请你找出并返回这两个正序数组的中位数。

进阶：你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗？

示例 1：

输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2
示例 2：

输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
示例 3：

输入：nums1 = [0,0], nums2 = [0,0]
输出：0.00000
示例 4：

输入：nums1 = [], nums2 = [1]
输出：1.00000
示例 5：

输入：nums1 = [2], nums2 = []
输出：2.00000

提示：

nums1.length == m
nums2.length == n
0 <= m <= 1000
0 <= n <= 1000
1 <= m + n <= 2000
-106 <= nums1[i], nums2[i] <= 106
通过次数307,367提交次数781,908

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	nums1 := []int{1, 2, 3}
	nums2 := []int{4, 5, 6}
	findMedianSortedArrays(nums1, nums2)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	ret := make([]int, 0)
	var i, j int
	for i = 0; i < len(nums1); i++ {
		for j = 0; j < len(nums2); j++ {
			if max(nums1[i], nums2[j]) {
				ret = append(ret, nums2[j])
			} else {
				ret = append(ret, nums1[i])
				break
			}
		}
	}
	return 1
}

func max(x, y int) bool {
	if x > y {
		return true
	}
	return false
}
