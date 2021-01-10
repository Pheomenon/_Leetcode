package main

/*
给定一个二叉树，返回它的 后序遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [3,2,1]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？

通过次数177,563提交次数240,439

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-postorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	var track func(node *TreeNode)
	track = func(node *TreeNode) {
		if node == nil {
			return
		}

		track(node.Left)
		track(node.Right)
		ans = append(ans, node.Val)
	}
	track(root)
	return ans
}
