package main

/*
给定一个 N 叉树，返回其节点值的前序遍历。
例如，给定一个3叉树:

返回其前序遍历: [1,3,5,6,2,4]。

说明:递归法很简单，你可以使用迭代法完成此题吗?

通过次数61,137提交次数82,591

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {

}

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	var track func(node *Node)
	ans := make([]int, 0)
	track = func(node *Node) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		for i := 0; i < len(node.Children); i++ {
			track(node.Children[i])
		}
	}
	track(root)
	return ans
}
