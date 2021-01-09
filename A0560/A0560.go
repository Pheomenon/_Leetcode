package main

import "fmt"

/*
给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。

示例 1 :
输入:nums = [1,1,1], k = 2
输出: 2 , [1,1] 与 [1,1] 为两种不同的情况。

说明 :
数组的长度为 [1, 20,000]。
数组中元素的范围是 [-1000, 1000] ，且整数 k 的范围是 [-1e7, 1e7]。
通过次数84,197提交次数187,423

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subarray-sum-equals-k
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	nums := []int{-1, -1, 1}
	k := 1
	fmt.Println(subarraySum(nums, k))
}

//只能正确处理只有正数的数组
func subarraySum(nums []int, k int) int {
	container := make([][]int, 0)
	ans := 0
	sum := nums[0]
	if sum == k {
		ans++
	}
	left, right := 0, 1
	for ; right < len(nums) || left != len(nums)-1; right++ {
		if right < len(nums) {
			sum += nums[right]
		}
		if sum == k {
			if right > len(nums)-1 {
				left++
				ans--
			}
			ans++
			container = append(container, []int{nums[left], nums[right]})
		} else if sum > k {
			for sum >= k {
				if sum == k {
					ans++
					container = append(container, []int{nums[left], nums[right]})
					break
				}
				sum -= nums[left]
				left++
			}
		}

	}
	return ans
}
