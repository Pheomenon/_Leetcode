package main

import (
	"fmt"
	"sort"
)

/*
给定一个无重复元素的数组candidates和一个目标数target，找出candidates中所有可以使数字和为target的组合。
candidates中的数字可以无限制重复被选取。

说明：

所有数字（包括target）都是正整数。
解集不能包含重复的组合。
示例 1：

输入：candidates = [2,3,6,7], target = 7,
所求解集为：
[
  [7],
  [2,2,3]
]
示例 2：

输入：candidates = [2,3,5], target = 8,
所求解集为：
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]

提示：

1 <= candidates.length <= 30
1 <= candidates[i] <= 200
candidate 中的每个元素都是独一无二的。
1 <= target <= 500
通过次数192,629提交次数268,320

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combination-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	candidates := []int{2, 3, 6, 7}
	fmt.Println(combinationSum(candidates, 7))
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	ans := make([][]int, 0)
	track := make([]int, 0)
	sum := 0

	var backtrack func(start int, track []int, sum int)
	backtrack = func(start int, track []int, sum int) {

		if sum >= target {
			if sum == target {
				tmp := make([]int, len(track))
				copy(tmp, track)
				ans = append(ans, tmp)
				return
			}
			return
		}

		for i := start; i < len(candidates); i++ {
			track = append(track, candidates[i])
			sum += candidates[i]
			backtrack(i, track, sum)
			sum -= track[len(track)-1]
			track = track[:len(track)-1]
		}
	}

	backtrack(0, track, sum)
	return ans
}
