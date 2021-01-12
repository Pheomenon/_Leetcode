package main

/*
给定一个 N 叉树，找到其最大深度。
最大深度是指从根节点到最远叶子节点的最长路径上的节点总数。
N 叉树输入按层序遍历序列化表示，每组子节点由空值分隔（请参见示例）。

示例 1：
输入：root = [1,null,3,2,4,null,5,6]
输出：3

示例 2：
输入：root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
输出：5

提示：

树的深度不会超过 1000 。
树的节点数目位于 [0, 104] 之间。
通过次数40,297提交次数56,660

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-depth-of-n-ary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {

}

type Node struct {
	Val      int
	Children []*Node
}

//DFS
func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	var track func(node *Node) int
	track = func(node *Node) int {
		depth := 0

		for i := 0; i < len(node.Children); i++ {
			depth = max(depth, track(node.Children[i]))
		}
		return depth + 1
	}
	ans := track(root)
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//BFS
//func maxDepth(root *Node) int {
//	if root == nil {
//		return 0
//	}
//	q := make([]*Node, 0)
//	depth := 0
//	q = append(q, root)
//
//	for len(q) > 0 {
//		sz := len(q)
//		for i := 0; i < sz; i++ {
//			node := q[0]
//			q = q[1:]
//
//			for _, v := range node.Children {
//				q = append(q, v)
//			}
//		}
//		depth++
//	}
//	return depth
//}
