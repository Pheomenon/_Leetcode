package main

/*
给定一个二叉树，判断它是否是高度平衡的二叉树。
本题中，一棵高度平衡二叉树定义为：
一个二叉树每个节点的左右两个子树的高度差的绝对值不超过 1 。

示例 1：
输入：root = [3,9,20,null,null,15,7]
输出：true

示例 2：
输入：root = [1,2,2,3,3,null,null,4,4]
输出：false

示例 3：
输入：root = []
输出：true

提示：

树中的节点数在范围 [0, 5000] 内
-104 <= Node.val <= 104
通过次数163,990提交次数297,658

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/balanced-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if !isBalanced(root.Left) || !isBalanced(root.Right) {
		return false
	}

	var track func(node *TreeNode) int
	track = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		depth := max(track(node.Left), track(node.Right)) + 1
		return depth
	}
	track(root)

	heightL := track(root.Left) + 1
	heightR := track(root.Right) + 1

	if abs(heightL-heightR) > 1 {
		return false
	}
	return true
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
