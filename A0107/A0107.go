package main

/*
给定一个二叉树，返回其节点值自底向上的层序遍历。
（即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
例如：
给定二叉树 [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
返回其自底向上的层序遍历为：

[
  [15,7],
  [9,20],
  [3]
]
通过次数118,495提交次数174,335

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	ans := make([][]int, 0)
	var track func(node *TreeNode)
	track = func(node *TreeNode) {
		for len(q) > 0 {
			sz := len(q)
			tmp := make([]int, 0)
			for i := 0; i < sz; i++ {
				node := q[0]
				q = q[1:]
				tmp = append(tmp, node.Val)
				if node.Left != nil {
					q = append(q, node.Left)
				}
				if node.Right != nil {
					q = append(q, node.Right)
				}
			}
			ans = append(ans, tmp)
		}
	}
	track(root)
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return ans
}
