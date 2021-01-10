package main

/*
给你一棵二叉树，请你返回满足以下条件的所有节点的值之和：

该节点的祖父节点的值为偶数。（一个节点的祖父节点是指该节点的父节点的父节点。）
如果不存在祖父节点值为偶数的节点，那么返回 0。

示例：
输入：root = [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]
输出：18
解释：图中红色节点的祖父节点的值为偶数，蓝色节点为这些红色节点的祖父节点。

提示：

树中节点的数目在 1 到10^4 之间。
每个节点的值在 1 到 100 之间。
通过次数9,501提交次数11,715

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sum-of-nodes-with-even-valued-grandparent
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumEvenGrandparent(root *TreeNode) int {
	ans := 0
	var track func(node *TreeNode)
	track = func(node *TreeNode) {
		if node == nil {
			return
		}

		ret := node.Val % 2
		if ret == 0 {
			if node.Left != nil {
				if node.Left.Left != nil {
					ans += node.Left.Left.Val
				}
				if node.Left.Right != nil {
					ans += node.Left.Right.Val
				}
			}
			if node.Right != nil {
				if node.Right.Left != nil {
					ans += node.Right.Left.Val
				}
				if node.Right.Right != nil {
					ans += node.Right.Right.Val
				}
			}
		}
		track(node.Left)
		track(node.Right)
	}
	track(root)
	return ans
}
