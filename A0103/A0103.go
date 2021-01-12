package main

/*
给定一个二叉树，返回其节点值的锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回锯齿形层序遍历如下：

[
  [3],
  [20,9],
  [15,7]
]
通过次数109,263提交次数191,696

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	ans := make([][]int, 0)
	q := make([]*TreeNode, 0)
	q = append(q, root)
	layer := 1

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
			if layer%2 != 1 {
				for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
					tmp[i], tmp[j] = tmp[j], tmp[i]
				}
			}
			layer++
			ans = append(ans, tmp)
		}
	}
	track(root)
	return ans
}
