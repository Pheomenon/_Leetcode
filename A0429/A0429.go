package main

/*
给定一个 N 叉树，返回其节点值的层序遍历。（即从左到右，逐层遍历）。
树的序列化输入是用层序遍历，每组子节点都由 null 值分隔（参见示例）。

示例 1：
输入：root = [1,null,3,2,4,null,5,6]
输出：[[1],[3,2,4],[5,6]]

示例 2：
输入：root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
输出：[[1],[2,3,4,5],[6,7,8,9,10],[11,12,13],[14]]

提示：

树的高度不会超过 1000
树的节点总数在 [0,10^4] 之间
通过次数36,767提交次数54,277

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {

}

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}

	ans := make([][]int, 0)
	q := make([]*Node, 0)
	q = append(q, root)

	for len(q) > 0 {
		sz := len(q)
		tmp := make([]int, 0)
		for i := 0; i < sz; i++ {
			node := q[0]
			q = q[1:]

			tmp = append(tmp, node.Val)
			for i := 0; i < len(node.Children); i++ {
				q = append(q, node.Children[i])
			}
		}
		ans = append(ans, tmp)
	}
	return ans
}
