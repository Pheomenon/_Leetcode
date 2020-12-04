package main

/*
给定一个二叉树，返回所有从根节点到叶子节点的路径。
说明: 叶子节点是指没有子节点的节点。

示例:
输入:
   1
 /   \
2     3
 \
  5

输出: ["1->2->5", "1->3"]

解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
通过次数86,462提交次数130,400

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-paths
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func binaryTreePaths(root *TreeNode) []string {
//	res := []string{}
//	dp(root)
//}
//
//func dp(root *TreeNode, str string) string {
//	if root == nil {
//		return str
//	}
//	str += "->"
//	str += strconv.Itoa(root.Val)
//	if root.Right != nil {
//		return dp(root.Right, str)
//	}
//	if root.Left != nil {
//		return dp(root.Left, str)
//	}
//	return str
//}
