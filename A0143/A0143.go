package main

/*
给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例 1:
给定链表 1->2->3->4, 重新排列为 1->4->2->3.

示例 2:
给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.
通过次数77,568提交次数130,340

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reorder-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	nums := []int{
		1, 2, 3, 4, 5,
	}
	head := linked(nums, 0)
	reorderList(head)
}

func linked(nums []int, idx int) *ListNode {
	if idx == len(nums) {
		return nil
	}
	return &ListNode{
		nums[idx],
		linked(nums, idx+1),
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	end := reverse(slow)
	start := head

	for start != nil { //如果链表长度是偶数，那么后半边链表将比前半边链表多1，当前半边指向nil时结束
		startNext := start.Next
		endPre := end.Next
		start.Next = end
		end.Next = startNext
		start = startNext
		end = endPre
	}
}

func reverse(head *ListNode) *ListNode {
	curr := head
	var prev *ListNode

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
