package main

import "fmt"

/*
给定一组不含重复元素的整数数组nums，返回该数组所有可能的子集（幂集）。
说明：解集不能包含重复的子集。

示例:
输入: nums = [1,2,3]
输出:
[
	[3],
 	[1],
 	[2],
 	[1,2,3],
	[1,3],
	[2,3],
	[1,2],
	[]
]
通过次数182,113提交次数229,408

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subsets
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(test(nums))
}

func test(nums []int) (ans [][]int) {
	ans = append(ans, []int{})
	for _, n := range nums {
		for _, v := range ans {
			ans = append(ans, append([]int{n}, v...))
		}
	}
	return
}

func subsets(nums []int) (ans [][]int) {
	set := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int(nil), set...))
			return
		}
		set = append(set, nums[cur])
		dfs(cur + 1)
		set = set[:len(set)-1]
		dfs(cur + 1)
	}
	dfs(0)
	return
}
