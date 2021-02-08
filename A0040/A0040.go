package main

import "sort"

/*
给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
candidates 中的每个数字在每个组合中只能使用一次。

说明：
所有数字（包括目标数）都是正整数。
解集不能包含重复的组合。
示例 1:

输入: candidates = [10,1,2,7,6,1,5], target = 8,
所求解集为:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
示例 2:

输入: candidates = [2,5,2,1,2], target = 5,
所求解集为:
[
 [1,2,2],
 [5]
]
通过次数130,406提交次数203,264

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combination-sum-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	candidates := []int{
		10, 1, 2, 7, 6, 1, 5,
	}
	target := 8
	combinationSum2(candidates, target)
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	ans := make([][]int, 0)
	var backtrack func(index, sum int, track []int)
	backtrack = func(index, sum int, track []int) {
		if sum >= target {
			if sum == target {
				tmp := make([]int, len(track))
				copy(tmp, track)
				ans = append(ans, tmp)
			}
			return
		}

		for i := index; i < len(candidates); i++ {
			if i-1 >= index && candidates[i-1] == candidates[i] {
				continue
			}
			track = append(track, candidates[i])
			backtrack(i+1, sum+candidates[i], track)
			track = track[:len(track)-1]
		}
	}
	backtrack(0, 0, []int{})
	return ans
}
