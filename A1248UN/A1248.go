package main

import "fmt"

/*
给你一个整数数组nums 和一个整数 k。

如果某个 连续 子数组中恰好有 k 个奇数数字，我们就认为这个子数组是「优美子数组」。

请返回这个数组中「优美子数组」的数目。

示例 1：

输入：nums = [1,1,2,1,1], k = 3
输出：2
解释：包含 3 个奇数的子数组是 [1,1,2,1] 和 [1,2,1,1] 。
示例 2：

输入：nums = [2,4,6], k = 1
输出：0
解释：数列中不包含任何奇数，所以不存在优美子数组。
示例 3：

输入：nums = [2,2,2,1,2,2,1,2,2,2], k = 2
输出：16

提示：

1 <= nums.length <= 50000
1 <= nums[i] <= 10^5
1 <= k <= nums.length
通过次数26,876提交次数49,356

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/count-number-of-nice-subarrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	nums := []int{1, 1, 2, 1, 1}
	k := 3
	fmt.Println(numberOfSubarrays(nums, k))
}

func numberOfSubarrays(nums []int, k int) int {
	odd := make([]int, 0)
	tmp := 0
	res := 0
	odd = append(odd, 1)
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 != 0 {
			tmp++
			odd = append(odd, tmp)
		}
	}
	for _, v := range odd {
		if v == k {
			res++
		}
	}
	return res
}
