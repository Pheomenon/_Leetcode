package main

/*
给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？
示例:
输入: 3
输出: 5
解释:
给定 n = 3, 一共有 5 种不同结构的二叉搜索树:

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3
通过次数99,631提交次数143,962

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/unique-binary-search-trees
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	numTrees(3)
}

func numTrees(n int) int {
	memo := make([]int, n+1)
	return helper(n, &memo)
}
func helper(n int, memo *[]int) int {
	if n == 1 || n == 0 {
		return 1
	}
	if (*memo)[n] > 0 {
		return (*memo)[n]
	}
	count := 0
	for i := 0; i <= n-1; i++ {
		count += helper(i, memo) * helper(n-i-1, memo)
	}
	(*memo)[n] = count
	return count
}
