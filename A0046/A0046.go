package main

import "fmt"

/*
给定一个 没有重复 数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
通过次数235,002提交次数303,703

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	nums := []int{
		5, 4, 6, 2,
	}
	fmt.Println(permute(nums))
}

func permute(nums []int) [][]int {
	var res = make([][]int, 0)
	track := make([]int, 0)
	var b func(nums []int, track []int)
	b = func(nums []int, track []int) {
		//触发结束条件
		if len(track) == len(nums) {
			//这里要特别注意如果直接把track append给res的化传的是引用，
			//track后面值被修改的时候会影响到res，所以这里必须手动Copy出track的一个副本来
			temp := make([]int, len(track))
			copy(temp, track)
			res = append(res, temp)
			return
		}

		for i := 0; i < len(nums); i++ {
			//排除不合法的选择
			if contain(nums[i], track) {
				continue
			}
			track = append(track, nums[i])
			b(nums, track)
			track = track[:len(track)-1]
		}
	}
	b(nums, track)
	return res
}

func contain(num int, track []int) bool {
	for j := 0; j < len(track); j++ {
		if track[j] == num {
			return true
		}
	}
	return false
}
