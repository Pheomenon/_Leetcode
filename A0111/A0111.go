/**
给定一个二叉树，找出其最小深度。
最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
说明：叶子节点是指没有子节点的节点。

示例 1：
输入：root = [3,9,20,null,null,15,7]
输出：2

示例 2：
输入：root = [2,null,3,null,4,null,5,null,6]
输出：5

提示：
树中节点数的范围在 [0, 105] 内
-1000 <= Node.val <= 1000
通过次数152,885提交次数338,489

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-depth-of-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

//BFS Version
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	head := -1
	layer := 1
	queue = append(queue, root)
	head++

	for head >= 0 {
		sz := len(queue)
		for i := 0; i < sz; i++ {
			node := queue[0]
			queue = queue[1:]
			head--
			if node.Right == nil && node.Left == nil {
				return layer
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
				head++
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				head++
			}
		}
		layer++
	}
	return layer
}

//func minDepth(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	if root.Left == nil && root.Right == nil {
//		return 1
//	}
//	result := math.MaxInt64
//	if root.Left != nil {
//		result = min(result, minDepth(root.Left))
//	}
//	if root.Right != nil {
//		result = min(result, minDepth(root.Right))
//	}
//	return result + 1
//}
//
//func min(x, y int) int {
//	if x < y {
//		return x
//	}
//	return y
//}
