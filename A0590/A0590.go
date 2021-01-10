package main

/*
给定一个 N 叉树，返回其节点值的后序遍历。

例如，给定一个 3叉树 :

返回其后序遍历: [5,6,3,2,4,1].

说明:递归法很简单，你可以使用迭代法完成此题吗?

通过次数43,901提交次数58,432

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {

}

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	ans := make([]int, 0)
	var postOrder func(root *Node)
	if root == nil {
		return []int{}
	}
	postOrder = func(root *Node) {
		if root == nil {
			return
		}
		for _, child := range root.Children {
			postOrder(child)
		}
		ans = append(ans, root.Val)
	}
	postOrder(root)
	return ans
}
