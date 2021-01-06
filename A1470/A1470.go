package main

/*
给你一个数组 nums ，数组中有 2n 个元素，按 [x1,x2,...,xn,y1,y2,...,yn] 的格式排列。

请你将数组按 [x1,y1,x2,y2,...,xn,yn] 格式重新排列，返回重排后的数组。

示例 1：

输入：nums = [2,5,1,3,4,7], n = 3
输出：[2,3,5,4,1,7]
解释：由于 x1=2, x2=5, x3=1, y1=3, y2=4, y3=7 ，所以答案为 [2,3,5,4,1,7]
示例 2：

输入：nums = [1,2,3,4,4,3,2,1], n = 4
输出：[1,4,2,3,3,2,4,1]
示例 3：

输入：nums = [1,1,2,2], n = 2
输出：[1,2,1,2]

提示：

1 <= n <= 500
nums.length == 2n
1 <= nums[i] <= 10^3
通过次数38,003提交次数44,791

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shuffle-the-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	nums := []int{
		1, 2, 3, 4, 4, 3, 2, 1,
	}
	n := 4
	shuffle(nums, n)
}

func shuffle(nums []int, n int) []int {
	阴 := make([]int, n)
	阳 := make([]int, n)
	for i, j := 0, n; i < n; i, j = i+1, j+1 {
		阴[i] = nums[i]
		阳[i] = nums[j]
	}
	ans := make([]int, len(nums))
	for i := 0; i < len(ans); i += 2 {
		ans[i] = 阴[i/2]
	}
	for j := 1; j < len(ans); j += 2 {
		ans[j] = 阳[j/2]
	}
	return ans
}
