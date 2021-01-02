/**
给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。
你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
返回滑动窗口中的最大值。

进阶：
你能在线性时间复杂度内解决此题吗？

示例:
输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
输出: [3,3,5,5,6,7]

解释:
  滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7

提示：
1 <= nums.length <= 10^5
-10^4 <= nums[i] <= 10^4
1 <= k <= nums.length
通过次数85,808提交次数174,499

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sliding-window-maximum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	nums := []int{7, 2, 4}
	k := 2
	fmt.Println(maxSlidingWindow(nums, k))
}

var a []int

type hp struct {
	sort.IntSlice
}

func (h hp) Less(i, j int) bool {
	return a[h.IntSlice[i]] > a[h.IntSlice[j]]
}

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func maxSlidingWindow(nums []int, k int) []int {
	a = nums
	q := &hp{make([]int, k)}
	for i := 0; i < k; i++ {
		q.IntSlice[i] = i
	}
	heap.Init(q)

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q.IntSlice[0]]
	for i := k; i < n; i++ {
		heap.Push(q, i)
		for q.IntSlice[0] <= i-k {
			heap.Pop(q)
		}
		ans = append(ans, nums[q.IntSlice[0]])
	}
	return ans
}

//以下解法超时
//++++++++++++++++++++++++++++++++++++++++++++++++
//func maxSlidingWindow(nums []int, k int) []int {
//	result := make([]int, 0, k)
//	if len(nums) == 0 {
//		return result
//	}
//	if k == 1 {
//		return nums
//	}
//	max := math.MinInt64
//	left, maxIndex := 0, 0
//	for i := 0; i < k; i++ {
//		if max < nums[i] {
//			max = nums[i]
//		}
//	}
//	result = append(result, max)
//
//	for right := k; right < len(nums); right++ {
//		if left == maxIndex {
//			max, maxIndex = maxInRange(nums, left, right)
//		}
//		left++
//		if max < nums[right] {
//			max = nums[right]
//			maxIndex = right
//		}
//		result = append(result, max)
//
//	}
//	return result
//}
//
//func maxInRange(nums []int, left, right int) (int, int) {
//	left++
//	max := math.MinInt64
//	maxIndex := 0
//	for i := left; i <= right; i++ {
//		if max < nums[i] {
//			max = nums[i]
//			maxIndex = i
//		}
//	}
//	return max, maxIndex
//}
