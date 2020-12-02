package main

import (
	"fmt"
	"sort"
)

/*
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。



示例：

给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
通过次数373,265提交次数1,237,912

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	sort.Sort(sort.IntSlice(nums))
	n := len(nums)
	res := make([][]int, 0)
	for i := 0; i < n; i++ {
		tuples := twoSum(nums, i+1, 0-nums[i])
		for _, v := range tuples {
			v = append(v, nums[i])
			res = append(res, v)
		}

		for i < n-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}

func twoSum(nums []int, start, target int) [][]int {
	ans := make([][]int, 0)
	if len(nums) == 0 {
		return nil
	}
	end := len(nums) - 1
	for start < end {
		left := nums[start]
		right := nums[end]
		sum := left + right
		if sum == target {
			ans = append(ans, []int{nums[start], nums[end]})
			for start < end && nums[end] == right {
				end--
			}
			for start < end && nums[start] == left {
				start++
			}
		} else if sum > target {
			for start < end && nums[end] == right {
				end--
			}
		} else if sum < target {
			for start < end && nums[start] == left {
				start++
			}
		}
	}
	return ans
}
