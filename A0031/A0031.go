package main

/*
实现获取 下一个排列 的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
必须 原地 修改，只允许使用额外常数空间。

示例 1：

输入：nums = [1,2,3]
输出：[1,3,2]
示例 2：

输入：nums = [3,2,1]
输出：[1,2,3]
示例 3：

输入：nums = [1,1,5]
输出：[1,5,1]
示例 4：

输入：nums = [1]
输出：[1]

提示：

1 <= nums.length <= 100
0 <= nums[i] <= 100
通过次数131,503提交次数361,658

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/next-permutation
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	nums := []int{5, 1, 1}
	nextPermutation(nums)
}

func nextPermutation(nums []int) {
	left := 0
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			left = i - 1
			break
		}
	}
	right := 0
	for i := len(nums) - 1; i > -1; i-- {
		if !min(nums[i], nums[left]) {
			right = i
			break
		}
	}
	if right != 0 || left != 0 {
		nums[left], nums[right] = nums[right], nums[left]
		reverse(nums, left+1)
	} else {
		reverse(nums, 0)
	}
}

func min(x, y int) bool {
	return x <= y
}

func reverse(nums []int, left int) {
	for i, j := left, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
