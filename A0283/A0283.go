/*
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

示例:
输入: [0,1,0,3,12]
输出: [1,3,12,0,0]

说明:
必须在原数组上操作，不能拷贝额外的数组。
尽量减少操作次数。
通过次数249,310提交次数395,094

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/move-zeroes
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "fmt"

func main() {
	nums := []int{0}
	moveZeroes(nums)
	for i := 0; i < len(nums); i++ {
		fmt.Print(nums[i])
	}
}

func moveZeroes(nums []int) {
	n := len(nums)
	times := 0
	if n == 0 || n == 1 {
		return
	}
	for i := 0; i < n; {
		if nums[i] == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, 0)
		}
		if nums[i] != 0 {
			i++
		}
		times++
		if times == n {
			return
		}
	}
}
