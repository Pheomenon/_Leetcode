package main

import "fmt"

/*
在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

示例 1:
输入: [3,2,1,5,6,4] 和 k = 2
输出: 5

示例 2:
输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
输出: 4
说明:

你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。

通过次数254,807提交次数393,645

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/kth-largest-element-in-an-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	nums := []int{5, 2, 4, 1, 3, 6, 0}
	k := 4
	fmt.Println(findKthLargest(nums, k))
}

func findKthLargest(nums []int, k int) int {
	target := len(nums) - k

	var partition func(left, right int) int
	partition = func(left, right int) int {
		pivot := nums[left]
		j := left
		for i := left + 1; i <= right; i++ {
			if nums[i] < pivot {
				j++
				nums[j], nums[i] = nums[i], nums[j]
			}
		}
		nums[j], nums[left] = nums[left], nums[j]
		return j
	}

	left, right := 0, len(nums)-1

	for {
		index := partition(left, right)
		if index == target {
			return nums[index]
		} else if index > target {
			right = index - 1
		} else {
			left = index + 1
		}
	}
}
