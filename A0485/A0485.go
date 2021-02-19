package main

import "fmt"

/*
给定一个二进制数组， 计算其中最大连续 1 的个数。
示例：

输入：[1,1,0,1,1,1]
输出：3
解释：开头的两位和最后的三位都是连续 1 ，所以最大连续 1 的个数是 3.
提示：
输入的数组只包含 0 和 1 。
输入数组的长度是正整数，且不超过 10,000。
通过次数92,228提交次数153,881

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/max-consecutive-ones
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	nums := []int{
		1, 1, 0, 1,
	}
	fmt.Println(findMaxConsecutiveOnes(nums))
}

func findMaxConsecutiveOnes(nums []int) int {
	left := -1
	right := 0
	ans := 0
	for right < len(nums) {
		if nums[right] == 1 {
			right++
		} else {
			ans = max(right-left-1, ans)
			left = right
			right++
		}
	}
	ans = max(right-left-1, ans)
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
