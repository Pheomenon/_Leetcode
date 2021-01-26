package main

import (
	"sort"
)

/*
给出一个区间的集合，请合并所有重叠的区间。

示例 1:
输入: intervals = [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

示例 2:
输入: intervals = [[1,4],[4,5]]
输出: [[1,5]]
解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
注意：输入类型已于2019年4月15日更改。 请重置默认代码定义以获取新方法签名。

提示：

intervals[i][0] <= intervals[i][1]
通过次数183,539提交次数416,267

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-intervals
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	intervals := [][]int{
		{1, 3}, {2, 6}, {8, 10}, {15, 18},
		//{1, 4}, {2, 3},
	}
	merge(intervals)
}

func merge(intervals [][]int) [][]int {
	ans := make([][]int, 0)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	prev := intervals[0]
	for i := 1; i < len(intervals); i++ {
		curr := intervals[i]
		if curr[0] > prev[1] {
			ans = append(ans, prev)
			prev = curr
		} else if curr[1] > prev[1] {
			prev[1] = curr[1]
		}
	}
	ans = append(ans, prev)
	return ans
}
