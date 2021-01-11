package main

/*
给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。
本题中，一个高度平衡二叉树是指一个二叉树每个节点的左右两个子树的高度差的绝对值不超过 1。
示例:
给定的有序链表： [-10, -3, 0, 5, 9],
一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：
      0
     / \
   -3   9
   /   /
 -10  5
通过次数69,407提交次数91,158

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	nums := []int{
		-10, -3, 0, 5, 9,
	}
	var newLinkList func(nums []int, idx int) *ListNode
	newLinkList = func(nums []int, idx int) *ListNode {
		if idx > len(nums)-1 {
			return nil
		}

		node := &ListNode{
			Val:  nums[idx],
			Next: newLinkList(nums, idx+1),
		}
		return node

	}
	linkList := newLinkList(nums, 0)
	sortedListToBST(linkList)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	//先将链表转成数组，然后从数组中间节点开始建树
	//curr := head
	//nums := make([]int,0)
	//for curr != nil{
	//	nums = append(nums, curr.Val)
	//	curr = curr.Next
	//}
	//if len(nums) == 0 {
	//	return nil
	//}
	//var track func(nums []int) *TreeNode
	//track = func(nums []int) *TreeNode {
	//	mid := len(nums) / 2
	//	if len(nums) == 0 {
	//		return nil
	//	}
	//
	//	node := &TreeNode{
	//		Val:   nums[mid],
	//		Left:  nil,
	//		Right: nil,
	//	}
	//	if mid == 0 {
	//		return node
	//	}
	//	node.Left = track(nums[:mid])
	//	node.Right = track(nums[mid+1:])
	//	return node
	//}
	//return track(nums)

	//直接建树
	var track func(head *ListNode, tail *ListNode) *TreeNode
	track = func(head *ListNode, tail *ListNode) *TreeNode {
		if head == tail {
			return nil
		}
		midNode := getMid(head, tail)

		node := &TreeNode{
			Val:   midNode.Val,
			Left:  nil,
			Right: nil,
		}
		node.Left = track(head, midNode)
		node.Right = track(midNode.Next, tail)
		return node
	}
	root := track(head, nil)
	return root
}

func getMid(head *ListNode, tail *ListNode) *ListNode {
	fast, slow := head, head
	for fast != tail && fast.Next != tail {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
