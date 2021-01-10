package main

import "fmt"

/*
有 n 个气球，编号为0 到 n-1，每个气球上都标有一个数字，这些数字存在数组 nums 中。
现在要求你戳破所有的气球。如果你戳破气球 i ，就可以获得 nums[left] * nums[i] * nums[right] 个硬币。
这里的 left 和 right 代表和 i 相邻的两个气球的序号。注意当你戳破了气球 i 后，气球 left 和气球 right 就变成了相邻的气球。
求所能获得硬币的最大数量。

说明:
你可以假设 nums[-1] = nums[n] = 1，但注意它们不是真实存在的所以并不能被戳破。
0 ≤ n ≤ 500, 0 ≤ nums[i] ≤ 100

示例:
输入: [3,1,5,8]
输出: 167
解释: nums = [3,1,5,8] --> [3,5,8] -->   [3,8]   -->  [8]  --> []
	coins =  3*1*5      +  3*5*8    +  1*3*8      + 1*8*1   = 167
通过次数36,503提交次数54,149

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/burst-balloons
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	nums := []int{3, 1, 5, 8}
	ans := maxCoins(nums)
	fmt.Println(ans)
}

func maxCoins(nums []int) int {
	ans := 0
	nums = append([]int{1}, nums...)
	nums = append(nums, 1)
	n := len(nums)
	var backtrack func(coins int, track, nums []int)
	backtrack = func(coins int, track, nums []int) {
		if len(track) == n-2 {
			ans = max(ans, coins)
			return
		}

		for i := 1; i < len(nums)-1; i++ {
			track = append(track, nums[i])
			coins = nums[i-1]*nums[i]*nums[i+1] + coins
			tmp := append(nums[:i], nums[i+1:]...)
			backtrack(coins, track, tmp)
		}
	}
	backtrack(0, []int{}, nums)
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
