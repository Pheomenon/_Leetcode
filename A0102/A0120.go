package main

/*
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。

示例：
二叉树：[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层序遍历结果：

[
  [3],
  [9,20],
  [15,7]
]
通过次数238,433提交次数372,766

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
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
	return ans
}
