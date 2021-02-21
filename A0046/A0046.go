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
	ans := make([][]int, 0)
	visited := map[int]bool{}
	var backtrack func(track []int, index int)
	backtrack = func(track []int, index int) {
		if len(track) == len(nums) {
			tmp := make([]int, len(track))
			copy(tmp, track)
			ans = append(ans, tmp)
			return
		}

		for i := index; i < len(nums); i++ {
			if visited[i] {
				continue
			}
			track = append(track, nums[i])
			visited[i] = true
			backtrack(track, index)
			track = track[:len(track)-1]
			visited[i] = false
		}
	}
	backtrack([]int{}, 0)
	return ans
}
