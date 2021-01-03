package main

import "fmt"

/*
给你一个链表和一个特定值 x ，请你对链表进行分隔，使得所有小于 x 的节点都出现在大于或等于 x 的节点之前。

你应当保留两个分区中每个节点的初始相对位置。

示例：

输入：head = 1->4->3->2->5->2, x = 3
输出：1->2->2->4->3->5
通过次数68,339提交次数111,763

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/partition-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	nums := []int{1, 4, 3, 2, 5, 2}
	head := &ListNode{
		Val:  nums[0],
		Next: nil,
	}
	curr := head
	for i := 1; i < 6; i++ {
		newNode := &ListNode{}
		newNode.Val = nums[i]
		curr.Next = newNode
		curr = newNode
	}
	fmt.Println(partition(head, 3))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	front, back := &ListNode{}, &ListNode{}
	curr := head
	backStart := back
	frontStart := front

	for curr != nil {
		if curr.Val < x {
			nextFront := &ListNode{
				Val:  curr.Val,
				Next: nil,
			}
			front.Next = nextFront
			front = nextFront
		} else {
			nextBack := &ListNode{
				Val:  curr.Val,
				Next: nil,
			}
			back.Next = nextBack
			back = nextBack
		}
		curr = curr.Next
	}
	front.Next = backStart.Next

	return frontStart.Next
}
