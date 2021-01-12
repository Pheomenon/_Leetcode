package main

/*
给定一个非空二叉树, 返回一个由每层节点平均值组成的数组。

示例 1：

输入：
    3
   / \
  9  20
    /  \
   15   7
输出：[3, 14.5, 11]
解释：
第 0 层的平均值是 3 ,
第1层是 14.5 ,
第2层是 11 。
因此返回 [3, 14.5, 11] 。

提示：

节点值的范围在32位有符号整数范围内。
通过次数53,155提交次数77,271

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/average-of-levels-in-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}
	ans := make([]float64, 0)
	q := make([]*TreeNode, 0)
	q = append(q, root)

	for len(q) > 0 {
		sz := len(q)
		var sum int
		for i := 0; i < sz; i++ {
			node := q[0]
			sum += node.Val
			q = q[1:]

			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		avg := avg(sum, sz)
		ans = append(ans, avg)
	}
	return ans
}

func avg(x, y int) float64 {
	return float64(x) / float64(y)
}
