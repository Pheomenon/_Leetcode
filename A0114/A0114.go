package main

/*
给定一个二叉树，原地将它展开为一个单链表。

例如，给定二叉树
    1
   / \
  2   5
 / \   \
3   4   6

将其展开为：
1
 \
  2
   \
    3
     \
      4
       \
        5
         \
          6
通过次数101,465提交次数141,622

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	list := preorder(root)
	for i := 1; i < len(list); i++ {
		pre, cur := list[i-1], list[i]
		pre.Left, pre.Right = nil, cur
	}
}

func preorder(node *TreeNode) []*TreeNode {
	list := []*TreeNode{}

	if node != nil {
		list = append(list, node)
		list = append(list, preorder(node.Left)...)
		list = append(list, preorder(node.Right)...)
	}
	return list
}
