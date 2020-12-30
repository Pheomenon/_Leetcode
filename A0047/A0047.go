package main

import (
	"fmt"
	"sort"
)

/*
给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。

示例 1：

输入：nums = [1,1,2]
输出：
[
	[1,1,2],
 	[1,2,1],
 	[2,1,1]
]

示例 2：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

提示：

1 <= nums.length <= 8
-10 <= nums[i] <= 10
通过次数128,417提交次数205,675

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	nums := []int{
		1, 1, 2,
	}
	fmt.Println(permuteUnique(nums))
}

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	track := make([]int, 0)
	vis := make([]bool, len(nums))

	var backtrack func(nums, track []int)
	backtrack = func(nums, track []int) {
		if len(nums) == len(track) {
			temp := make([]int, len(track))
			copy(temp, track)
			ans = append(ans, temp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if vis[i] || i > 0 && !vis[i-1] && nums[i] == nums[i-1] {
				continue
			}

			track = append(track, nums[i])
			vis[i] = true
			backtrack(nums, track)
			vis[i] = false
			track = track[:len(track)-1]
		}
	}
	backtrack(nums, track)
	return ans
}
